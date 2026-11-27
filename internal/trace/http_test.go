package trace

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHeaders(t *testing.T) {
	tests := []struct {
		name, request, trace, session string
		valid                         bool
	}{
		{"valid", "request_123", "trace-123", "session_123", true},
		{"missing", "", "", "", false},
		{"short", "short", "short", "short", false},
		{"long", strings.Repeat("x", 65), strings.Repeat("y", 65), strings.Repeat("z", 65), false},
		{"injection", "abc12345\r\nX: secret", "abc12345\nsecret", "abc12345 secret", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := make(http.Header)
			h.Set(HeaderRequestID, tt.request)
			h.Set(HeaderTraceID, tt.trace)
			h.Set(HeaderSessionID, tt.session)
			ids := IDsFromHeaders(h)
			if !IsValidID(ids.RequestID) || !IsValidID(ids.TraceID) {
				t.Fatalf("invalid result: %+v", ids)
			}
			if tt.valid {
				if ids.RequestID != tt.request || ids.TraceID != tt.trace || ids.SessionID != tt.session {
					t.Fatal(ids)
				}
			} else if ids.RequestID == tt.request || ids.TraceID != ids.RequestID || ids.SessionID != "" {
				t.Fatal(ids)
			}
		})
	}
	h := http.Header{}
	h.Add(HeaderRequestID, "request1")
	h.Add(HeaderRequestID, "request2")
	if ids := IDsFromHeaders(h); ids.RequestID == "request1" || ids.RequestID == "request2" {
		t.Fatal(ids)
	}
}

func TestMiddlewareResponsesAndSafeLogs(t *testing.T) {
	var buf bytes.Buffer
	oldLog, oldAccess := slog.Default(), AccessLog
	t.Cleanup(func() { slog.SetDefault(oldLog); AccessLog = oldAccess })
	slog.SetDefault(slog.New(slog.NewJSONHandler(&buf, nil)))
	AccessLog = slog.New(slog.NewJSONHandler(&buf, nil))
	for _, mode := range []string{"success", "error", "panic", "early"} {
		t.Run(mode, func(t *testing.T) {
			buf.Reset()
			h := Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if mode != "early" {
					StateFromContext(r.Context()).Route = "/customers/:id"
				}
				switch mode {
				case "error":
					RespondError(w, r.Context(), http.StatusBadRequest, "bad_request", "bad request")
				case "panic":
					panic("secret-panic-value")
				case "early":
					w.WriteHeader(http.StatusForbidden)
				default:
					w.Write([]byte("ok"))
				}
			}))
			r := httptest.NewRequest("GET", "/customers/private-name?secret=query-value", nil)
			r.Header.Set(HeaderRequestID, "request_123")
			r.Header.Set(HeaderTraceID, "trace_1234")
			r.Header.Set(HeaderSessionID, "session_123")
			r.Header.Set("Authorization", "Bearer secret-token")
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)
			if w.Header().Get(HeaderRequestID) != "request_123" || w.Header().Get(HeaderTraceID) != "trace_1234" {
				t.Fatal(w.Header())
			}
			if mode == "panic" || mode == "error" {
				var body ErrorBody
				if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil || body.Error.RequestID != "request_123" {
					t.Fatalf("%s %v", w.Body, err)
				}
				if mode == "panic" && (w.Code != 500 || body.Error.Code != "internal_error") {
					t.Fatal(w.Code, body)
				}
			}
			lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
			access := 0
			for _, l := range lines {
				var m map[string]interface{}
				if err := json.Unmarshal([]byte(l), &m); err != nil {
					t.Fatal(err)
				}
				if m["service"] != Service || m["request_id"] != "request_123" || m["trace_id"] != "trace_1234" || m["session_id"] != "session_123" {
					t.Fatal(m)
				}
				if m["msg"] == "access" {
					access++
				}
			}
			if access != 1 {
				t.Fatalf("access=%d logs=%s", access, buf.String())
			}
			for _, secret := range []string{"private-name", "query-value", "secret-token", "secret-panic-value"} {
				if strings.Contains(buf.String(), secret) {
					t.Fatalf("sensitive data logged: %q", secret)
				}
			}
		})
	}
}

func TestMiddlewareGeneratesAndSanitizesResponseIDs(t *testing.T) {
	var buf bytes.Buffer
	old := AccessLog
	AccessLog = slog.New(slog.NewJSONHandler(&buf, nil))
	t.Cleanup(func() { AccessLog = old })
	for _, value := range []string{"", "short", "secret\r\nInjected"} {
		buf.Reset()
		r := httptest.NewRequest("GET", "/", nil)
		if value != "" {
			r.Header.Set(HeaderRequestID, value)
			r.Header.Set(HeaderTraceID, value)
		}
		w := httptest.NewRecorder()
		Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("ok")) })).ServeHTTP(w, r)
		requestID, traceID := w.Header().Get(HeaderRequestID), w.Header().Get(HeaderTraceID)
		if !IsValidID(requestID) || requestID != traceID || requestID == value {
			t.Fatalf("request=%q trace=%q", requestID, traceID)
		}
		if value != "" && strings.Contains(buf.String(), value) {
			t.Fatal("invalid header logged")
		}
	}
}
