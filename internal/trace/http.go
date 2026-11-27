package trace

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

// AccessLog is separate so access events are emitted even above INFO.
var AccessLog = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type stateKey struct{}
type State struct{ Route string }

func StateFromContext(ctx context.Context) *State { s, _ := ctx.Value(stateKey{}).(*State); return s }

// IDsFromHeaders validates each complete value. Ambiguous repeated headers are rejected.
func IDsFromHeaders(h http.Header) IDs {
	one := func(name string) string {
		v := h.Values(name)
		if len(v) != 1 {
			return ""
		}
		return SanitizeHeader(name, v[0])
	}
	ids := IDs{SessionID: one(HeaderSessionID), TraceID: one(HeaderTraceID), RequestID: one(HeaderRequestID)}
	if ids.RequestID == "" {
		ids.RequestID = NewID()
	}
	if ids.TraceID == "" {
		ids.TraceID = ids.RequestID
	}
	return ids
}

type recorder struct {
	http.ResponseWriter
	status, bytes int
}

func (w *recorder) WriteHeader(code int) {
	if w.status != 0 {
		return
	}
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}
func (w *recorder) Write(p []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	n, err := w.ResponseWriter.Write(p)
	w.bytes += n
	return n, err
}
func (w *recorder) Unwrap() http.ResponseWriter { return w.ResponseWriter }

// RespondError writes the gateway-compatible error envelope.
func RespondError(w http.ResponseWriter, ctx context.Context, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(StandardError(code, message, FromContext(ctx).RequestID))
}

// Middleware wraps all requests, including ones which never reach BeforeRouter.
// To avoid logging path parameters, path and route are route templates, not URLs.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ids := IDsFromHeaders(r.Header)
		st := &State{}
		ctx := context.WithValue(WithIDs(r.Context(), ids), stateKey{}, st)
		r = r.WithContext(ctx)
		w.Header().Set(HeaderRequestID, ids.RequestID)
		w.Header().Set(HeaderTraceID, ids.TraceID)
		rec := &recorder{ResponseWriter: w}
		defer func() {
			if p := recover(); p != nil {
				// Panic values can contain user data and must not be logged.
				Logger(ctx, Service).Error("panic recovered", "stack", string(debug.Stack()))
				if rec.status == 0 {
					RespondError(rec, ctx, http.StatusInternalServerError, "internal_error", "internal server error")
				}
			}
			status := rec.status
			if status == 0 {
				status = http.StatusOK
			}
			route := st.Route
			if route == "" {
				route = "<unmatched>"
			}
			// Access lines must not disappear when application log level is warn/error.
			accessLogger(ctx).InfoContext(ctx, "access", "method", safeMethod(r.Method), "path", route, "route", route, "status", status, "duration_ms", time.Since(start).Milliseconds(), "bytes", rec.bytes)
		}()
		next.ServeHTTP(rec, r)
	})
}
func accessLogger(ctx context.Context) *slog.Logger {
	ids := FromContext(ctx)
	return AccessLog.With("service", Service, "session_id", ids.SessionID, "trace_id", ids.TraceID, "request_id", ids.RequestID)
}

func safeMethod(m string) string {
	switch m {
	case "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS":
		return m
	}
	return "OTHER"
}
