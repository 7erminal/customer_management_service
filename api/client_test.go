package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"customer_management_service/tracing"
)

// TestOutboundForwardsSameIDs verifies that an outbound call carries the
// SAME tracing identifiers from the request context, and no new request
// ID is minted.
func TestOutboundForwardsSameIDs(t *testing.T) {
	var got http.Header
	var gotBody map[string]string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Clone()
		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &gotBody)
	}))
	defer srv.Close()

	state := &tracing.State{SessionID: "session_123456", TraceID: "trace_12345678", RequestID: "request_12345678"}
	ctx := tracing.NewContext(context.Background(), state)

	request := NewRequest(srv.URL, "/v1/accounts", POST)
	request.Params["UserId"] = "42"
	client := Client{Request: request, Type_: "body"}

	res, err := client.SendRequestWithContext(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if got.Get(tracing.HeaderRequestID) != "request_12345678" {
		t.Errorf("outbound X-Request-Id = %q, want the same inbound value", got.Get(tracing.HeaderRequestID))
	}
	if got.Get(tracing.HeaderTraceID) != "trace_12345678" {
		t.Errorf("outbound X-Trace-Id = %q", got.Get(tracing.HeaderTraceID))
	}
	if got.Get(tracing.HeaderSessionID) != "session_123456" {
		t.Errorf("outbound X-Session-Id = %q", got.Get(tracing.HeaderSessionID))
	}
	if gotBody["UserId"] != "42" {
		t.Errorf("body params not sent: %v", gotBody)
	}
}

// TestOutboundWithoutContext verifies requests still work without a
// tracing context (background jobs, direct calls).
func TestOutboundWithoutContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer srv.Close()

	request := NewRequest(srv.URL, "/v1/accounts", GET)
	client := Client{Request: request, Type_: "params"}

	res, err := client.SendRequest()
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if res.StatusCode != 200 {
		t.Fatalf("status = %d", res.StatusCode)
	}
}

// TestOutboundHeadersAreSanitizedSameValues: identifiers are forwarded
// verbatim (they were already validated inbound); the outbound helper
// must not invent or mutate them.
func TestOutboundHeadersAreSameValues(t *testing.T) {
	state := &tracing.State{SessionID: "", TraceID: "trace_12345678", RequestID: "request_12345678"}
	h := http.Header{}
	tracing.ApplyOutboundHeaders(tracing.NewContext(context.Background(), state), h)
	if _, ok := h[http.CanonicalHeaderKey(tracing.HeaderSessionID)]; ok {
		t.Error("empty session ID must not be forwarded")
	}
	if h.Get(tracing.HeaderRequestID) != "request_12345678" || h.Get(tracing.HeaderTraceID) != "trace_12345678" {
		t.Errorf("outbound headers = %v", h)
	}
}
