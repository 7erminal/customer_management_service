// Package middleware implements gateway tracing for inbound requests:
// header sanitisation, context propagation, response headers, panic
// recovery with JSON errors and a single structured access-log line per
// request.
package middleware

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"customer_management_service/logging"
	"customer_management_service/tracing"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// Trace is the outer http middleware. It must wrap the beego handler
// (use beego.RunWithMiddleWares("", middleware.Trace)) so that it also
// covers panics, early error responses (404/405/413) and static files.
func Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		state := &tracing.State{}
		state.SessionID = tracing.SanitizeID(r.Header.Get(tracing.HeaderSessionID))
		state.TraceID = tracing.SanitizeID(r.Header.Get(tracing.HeaderTraceID))
		state.RequestID = tracing.SanitizeID(r.Header.Get(tracing.HeaderRequestID))
		// Missing request ID -> generate one; missing trace ID -> use
		// the request ID. Invalid values are discarded by SanitizeID
		// and never logged.
		if state.RequestID == "" {
			state.RequestID = tracing.NewRequestID()
		}
		if state.TraceID == "" {
			state.TraceID = state.RequestID
		}

		r = r.WithContext(tracing.NewContext(r.Context(), state))

		rec := &recordingResponseWriter{ResponseWriter: w, status: http.StatusOK}
		// Set the response headers before any write so they are present
		// on success, error, early responses and panics alike.
		h := rec.Header()
		h.Set(tracing.HeaderRequestID, state.RequestID)
		h.Set(tracing.HeaderTraceID, state.TraceID)

		defer func() {
			pr := recover()
			if pr != nil && !isIntentionalAbort(pr) {
				stack := debug.Stack()
				logging.For(r.Context()).Error("panic recovered",
					"error", fmt.Sprint(pr),
					"stack", string(stack),
					"method", r.Method,
					"path", r.URL.Path,
				)
				if !rec.wroteHeader {
					WriteError(rec, http.StatusInternalServerError, "internal_error",
						"internal server error", state.RequestID)
				}
			}
			logRequest(r, rec, state, time.Since(start))
		}()

		next.ServeHTTP(rec, r)
	})
}

// RoutePattern is a beego BeforeExec filter that records the matched
// route pattern so the access-log line can include it.
func RoutePattern() beego.FilterFunc {
	return func(ctx *context.Context) {
		if pattern, ok := ctx.Input.GetData("RouterPattern").(string); ok && pattern != "" {
			if state, ok := tracing.FromContext(ctx.Request.Context()); ok {
				state.RoutePattern = pattern
			}
		}
	}
}

// RespondError writes a JSON error envelope that always carries the
// request_id from the current request context:
//
//	{"error": {"code": "...", "message": "...", "request_id": "..."}}
func RespondError(c *beego.Controller, status int, code, message string) {
	requestID := tracing.RequestIDFromContext(c.Ctx.Request.Context())
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = ErrorBody(code, message, requestID)
	c.ServeJSON()
}

// WriteError writes the JSON error envelope to an http.ResponseWriter.
func WriteError(w http.ResponseWriter, status int, code, message, requestID string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorBody(code, message, requestID))
}

// ErrorBody builds the JSON error envelope.
func ErrorBody(code, message, requestID string) map[string]interface{} {
	return map[string]interface{}{
		"error": map[string]interface{}{
			"code":       code,
			"message":    message,
			"request_id": requestID,
		},
	}
}

// logRequest emits the single access-log line for the request.
func logRequest(r *http.Request, rec *recordingResponseWriter, state *tracing.State, d time.Duration) {
	logging.For(r.Context()).Info("http_request",
		"method", r.Method,
		"path", r.URL.Path, // path only, query string values are never logged
		"route", state.RoutePattern,
		"status", rec.status,
		"duration_ms", float64(d.Microseconds())/1000.0,
		"bytes", rec.bytes,
	)
}

// isIntentionalAbort reports whether the recovered value is one of the
// framework's request-abort sentinels rather than a real panic.
func isIntentionalAbort(pr interface{}) bool {
	return pr == http.ErrAbortHandler || pr == beego.ErrAbort
}

// recordingResponseWriter records status code and response size for the
// access log while delegating to the real writer.
type recordingResponseWriter struct {
	http.ResponseWriter
	status      int
	bytes       int64
	wroteHeader bool
}

func (w *recordingResponseWriter) WriteHeader(code int) {
	if !w.wroteHeader {
		w.status = code
		w.wroteHeader = true
	}
	w.ResponseWriter.WriteHeader(code)
}

func (w *recordingResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.wroteHeader = true
	}
	n, err := w.ResponseWriter.Write(b)
	w.bytes += int64(n)
	return n, err
}

func (w *recordingResponseWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *recordingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, fmt.Errorf("recordingResponseWriter: underlying writer does not support Hijack")
	}
	return h.Hijack()
}
