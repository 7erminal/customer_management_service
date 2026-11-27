package main

import (
	"bytes"
	"customer_management_service/internal/trace"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestBeegoTracingLifecycle(t *testing.T) {
	configureTracing()
	var buf bytes.Buffer
	old, access := slog.Default(), trace.AccessLog
	t.Cleanup(func() { slog.SetDefault(old); trace.AccessLog = access })
	slog.SetDefault(slog.New(slog.NewJSONHandler(&buf, nil)))
	trace.AccessLog = slog.New(slog.NewJSONHandler(&buf, nil))
	beego.InsertFilter("/trace-early", beego.BeforeRouter, func(c *beecontext.Context) {
		c.Output.SetStatus(http.StatusForbidden)
		_ = c.Output.Body([]byte("blocked"))
	})
	beego.Get("/trace-test/:id", func(c *beecontext.Context) {
		switch c.Input.Param(":id") {
		case "panic":
			panic("private panic")
		case "error":
			trace.RespondError(c.ResponseWriter, c.Request.Context(), http.StatusBadRequest, "bad_request", "invalid")
		default:
			c.Output.Body([]byte("ok"))
		}
	})
	beego.BeeApp.Handlers.Init()
	handler := trace.Middleware(beego.BeeApp.Handlers)
	for _, tc := range []struct {
		path   string
		status int
		route  string
	}{
		{"/trace-test/success", 200, "/trace-test/:id"}, {"/trace-test/error", 400, "/trace-test/:id"}, {"/trace-test/panic", 500, "/trace-test/:id"}, {"/unmatched-private", 404, "<unmatched>"}, {"/trace-early", 403, "<unmatched>"},
	} {
		t.Run(tc.path, func(t *testing.T) {
			buf.Reset()
			r := httptest.NewRequest("GET", tc.path+"?token=querysecret", nil)
			r.Header.Set(trace.HeaderRequestID, "request_123")
			r.Header.Set(trace.HeaderTraceID, "trace_1234")
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, r)
			if w.Code != tc.status {
				t.Fatalf("status %d body %s", w.Code, w.Body.String())
			}
			if w.Header().Get(trace.HeaderRequestID) != "request_123" || w.Header().Get(trace.HeaderTraceID) != "trace_1234" {
				t.Fatal(w.Header())
			}
			if tc.status == 500 || tc.status == 400 {
				var body trace.ErrorBody
				if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil || body.Error.RequestID != "request_123" {
					t.Fatalf("body %s %v", w.Body.String(), err)
				}
			}
			var accessCount int
			for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
				var m map[string]interface{}
				if err := json.Unmarshal([]byte(line), &m); err != nil {
					t.Fatal(err)
				}
				if m["msg"] == "access" {
					accessCount++
					if m["route"] != tc.route {
						t.Fatalf("route %v, want %s", m["route"], tc.route)
					}
				}
			}
			if accessCount != 1 {
				t.Fatalf("access count %d logs %s", accessCount, buf.String())
			}
			if strings.Contains(buf.String(), "querysecret") || strings.Contains(buf.String(), "private panic") || strings.Contains(buf.String(), "unmatched-private") {
				t.Fatal(buf.String())
			}
		})
	}
}
