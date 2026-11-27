// Package trace contains the tracing primitives copied from the gateway's
// github.com/msys/go-common/trace (commit a79a370ee178066e3f5a163ce69f0813cb6b249c).
// Replace this local subset with the published shared module when available.
package trace

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"net/http"
	"regexp"
)

const (
	HeaderSessionID = "X-Session-Id"
	HeaderTraceID   = "X-Trace-Id"
	HeaderRequestID = "X-Request-Id"
	Service         = "customer_management_service"
)

type IDs struct{ SessionID, TraceID, RequestID string }
type contextKey string

const idsKey contextKey = "msys/trace:ids"

func WithIDs(ctx context.Context, ids IDs) context.Context {
	return context.WithValue(ctx, idsKey, ids)
}
func FromContext(ctx context.Context) IDs {
	if ctx != nil {
		if ids, ok := ctx.Value(idsKey).(IDs); ok {
			return ids
		}
	}
	return IDs{}
}

var validID = regexp.MustCompile(`^[A-Za-z0-9_-]{8,64}$`)

func IsValidID(s string) bool { return validID.MatchString(s) }
func SanitizeHeader(_ string, raw string) string {
	if IsValidID(raw) {
		return raw
	}
	return ""
}

// NewID uses 128 bits of cryptographically secure randomness, encoded as hex.
func NewID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic("unable to generate request id")
	}
	return hex.EncodeToString(b[:])
}
func ApplyToHeader(h http.Header, ids IDs) {
	if ids.RequestID != "" {
		h.Set(HeaderRequestID, ids.RequestID)
	} else {
		h.Del(HeaderRequestID)
	}
	if ids.TraceID != "" {
		h.Set(HeaderTraceID, ids.TraceID)
	} else {
		h.Del(HeaderTraceID)
	}
	if ids.SessionID != "" {
		h.Set(HeaderSessionID, ids.SessionID)
	} else {
		h.Del(HeaderSessionID)
	}
}
func Logger(ctx context.Context, service string) *slog.Logger {
	ids := FromContext(ctx)
	return slog.With("service", service, "session_id", ids.SessionID, "trace_id", ids.TraceID, "request_id", ids.RequestID)
}

type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}
type ErrorDetail struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func StandardError(code, message, requestID string) ErrorBody {
	return ErrorBody{Error: ErrorDetail{Code: code, Message: message, RequestID: requestID}}
}
