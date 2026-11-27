package tracing

import (
	"context"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestSanitizeID(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  string
	}{
		{"valid lowercase hex", "abc123def456", "abc123def456"},
		{"valid upper and dash", "Trace-ID-9", "Trace-ID-9"},
		{"valid underscore", "sess_12345678", "sess_12345678"},
		{"minimum length", "12345678", "12345678"},
		{"maximum length", "a60b7c8d9e0f1g2h3i4j5k6l7m8n9o0p1q2r3s4t5u6v7w8x9y0z1a2b3c4d5e6f", "a60b7c8d9e0f1g2h3i4j5k6l7m8n9o0p1q2r3s4t5u6v7w8x9y0z1a2b3c4d5e6f"},
		{"too short", "abc", ""},
		{"empty", "", ""},
		{"too long", "1234567890123456789012345678901234567890123456789012345678901234567890", ""},
		{"injection semicolon", "abc123;DROP TABLE users", ""},
		{"injection space", "abc123 def456", ""},
		{"injection quotes", "abc\"onerror=alert(1)", ""},
		{"newline", "abc123\nxyz", ""},
		{"carriage return", "abc123\r\nxyz", ""},
		{"tab", "abc\t12345", ""},
		{"comma joined", "id1,id2,id3", ""},
		{"pipe joined", "id1|id2", ""},
		{"unicode", "идентификатор", ""},
		{"control chars", "\x01\x02abc456", ""},
		{"valid with surrounding spaces", "  abc123def  ", "abc123def"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SanitizeID(tt.value); got != tt.want {
				t.Errorf("SanitizeID(%q) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestNewRequestID(t *testing.T) {
	for i := 0; i < 100; i++ {
		id := NewRequestID()
		if SanitizeID(id) != id {
			t.Fatalf("generated request ID does not pass its own sanitisation: %q", id)
		}
		if !hexIDPattern.MatchString(id) {
			t.Fatalf("unexpected request ID format: %q", id)
		}
	}
}

func TestContextRoundTrip(t *testing.T) {
	state := &State{SessionID: "session_1", TraceID: "trace_1", RequestID: "request_1"}
	ctx := NewContext(context.Background(), state)
	got, ok := FromContext(ctx)
	if !ok || got != state {
		t.Fatalf("expected stored state, got %v (ok=%v)", got, ok)
	}
	if RequestIDFromContext(ctx) != "request_1" {
		t.Fatalf("RequestIDFromContext = %q", RequestIDFromContext(ctx))
	}
	if RequestIDFromContext(context.Background()) != "" {
		t.Fatal("expected empty request ID for context without tracing state")
	}
}

func TestApplyOutboundHeaders(t *testing.T) {
	state := &State{SessionID: "session_1", TraceID: "trace_1", RequestID: "request_1"}
	h := http.Header{}
	ApplyOutboundHeaders(NewContext(context.Background(), state), h)
	if h.Get(HeaderRequestID) != "request_1" || h.Get(HeaderTraceID) != "trace_1" || h.Get(HeaderSessionID) != "session_1" {
		t.Fatalf("outbound headers = %v", h)
	}

	// No tracing state -> no headers set.
	h2 := http.Header{}
	ApplyOutboundHeaders(context.Background(), h2)
	if len(h2) != 0 {
		t.Fatalf("expected no outbound headers without tracing state, got %v", h2)
	}
}

func TestOutboundHeadersViaTestServer(t *testing.T) {
	var got http.Header
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
	}))
	defer srv.Close()

	state := &State{SessionID: "session_9", TraceID: "trace_9", RequestID: "request_9"}
	req := httptest.NewRequest("GET", srv.URL, nil).WithContext(NewContext(context.Background(), state))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	if got.Get(HeaderRequestID) != "request_9" || got.Get(HeaderTraceID) != "trace_9" || got.Get(HeaderSessionID) != "session_9" {
		t.Fatalf("forwarded headers = %v", got)
	}
}

var hexIDPattern = regexp.MustCompile(`^[0-9a-f]{32}$`)
