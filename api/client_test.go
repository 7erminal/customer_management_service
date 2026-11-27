package api

import (
	"context"
	"customer_management_service/internal/trace"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForwardTraceHeadersAndCancellation(t *testing.T) {
	ids := trace.IDs{SessionID: "session_123", TraceID: "trace_123", RequestID: "request_123"}
	called := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		for k, v := range map[string]string{trace.HeaderSessionID: ids.SessionID, trace.HeaderTraceID: ids.TraceID, trace.HeaderRequestID: ids.RequestID} {
			if r.Header.Get(k) != v {
				t.Errorf("%s: %q", k, r.Header.Get(k))
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	req := NewRequest(server.URL, "/", POST)
	req.HeaderField[trace.HeaderRequestID] = "different"
	client := Client{Request: req, Type_: "body"}
	ctx := trace.WithIDs(context.Background(), ids)
	res, err := client.SendRequest(ctx)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()
	if !called {
		t.Fatal("not called")
	}
	ctx, cancel := context.WithCancel(ctx)
	cancel()
	if _, err = client.SendRequest(ctx); err == nil {
		t.Fatal("cancelled request succeeded")
	}
}
