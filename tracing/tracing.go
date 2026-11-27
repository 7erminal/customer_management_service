// Package tracing carries gateway tracing identifiers (X-Session-Id,
// X-Trace-Id, X-Request-Id) through the request context, outgoing HTTP
// calls and log lines.
//
// NOTE: this is a service-local copy of the gateway trace contract
// (header names, sanitisation rules). Once the shared trace module from
// the gateway work is published, this package should be replaced by it
// keeping only the same small surface.
package tracing

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
)

// Headers set by the API gateway.
const (
	HeaderSessionID = "X-Session-Id"
	HeaderTraceID   = "X-Trace-Id"
	HeaderRequestID = "X-Request-Id"
)

// ID length bounds enforced on every inbound identifier.
const (
	MinIDLength = 8
	MaxIDLength = 64
)

type contextKey struct{}

// State holds the tracing identifiers for one request. It is stored in
// the request context as a pointer so the access-log middleware can see
// values (e.g. the matched route pattern) attached later by filters.
type State struct {
	SessionID    string
	TraceID      string
	RequestID    string
	RoutePattern string
}

// SanitizeID validates a client supplied identifier. Valid IDs are
// 8-64 characters from [A-Za-z0-9_-]. Anything else (empty, wrong
// length, whitespace, quotes, control characters, injection attempts)
// is discarded: the return value is "" and the raw value is never
// logged.
func SanitizeID(value string) string {
	value = strings.TrimSpace(value)
	if len(value) < MinIDLength || len(value) > MaxIDLength {
		return ""
	}
	for i := 0; i < len(value); i++ {
		c := value[i]
		switch {
		case c >= 'A' && c <= 'Z':
		case c >= 'a' && c <= 'z':
		case c >= '0' && c <= '9':
		case c == '-' || c == '_':
		default:
			return ""
		}
	}
	return value
}

// NewRequestID generates a request identifier that always passes
// SanitizeID (32 hex characters).
func NewRequestID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		// crypto/rand failing is fatal for ID generation; fall back to
		// a fixed-format placeholder rather than panicking.
		return "req-00000000000000000000000000000000"
	}
	return hex.EncodeToString(b)
}

// NewContext returns a copy of ctx carrying state.
func NewContext(ctx context.Context, state *State) context.Context {
	return context.WithValue(ctx, contextKey{}, state)
}

// FromContext extracts the tracing state, if any.
func FromContext(ctx context.Context) (*State, bool) {
	if ctx == nil {
		return nil, false
	}
	state, ok := ctx.Value(contextKey{}).(*State)
	return state, ok
}

// RequestIDFromContext returns the current request ID, or "" when the
// context carries no tracing state.
func RequestIDFromContext(ctx context.Context) string {
	if state, ok := FromContext(ctx); ok {
		return state.RequestID
	}
	return ""
}

// ApplyOutboundHeaders copies the tracing identifiers onto the headers
// of an outgoing request. It always forwards the SAME identifiers that
// were received from the gateway (or generated for this request); new
// request IDs are never minted for outbound calls.
func ApplyOutboundHeaders(ctx context.Context, h http.Header) {
	state, ok := FromContext(ctx)
	if !ok || state == nil {
		return
	}
	if state.RequestID != "" {
		h.Set(HeaderRequestID, state.RequestID)
	}
	if state.TraceID != "" {
		h.Set(HeaderTraceID, state.TraceID)
	}
	if state.SessionID != "" {
		h.Set(HeaderSessionID, state.SessionID)
	}
}
