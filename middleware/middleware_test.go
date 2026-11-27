package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"customer_management_service/logging"
	"customer_management_service/tracing"

	beego "github.com/beego/beego/v2/server/web"
)

// echoController writes the tracing IDs it sees in its request context
// back into the response body, so tests can verify context propagation.
type echoController struct {
	beego.Controller
}

func (c *echoController) Get() {
	state, _ := tracing.FromContext(c.Ctx.Request.Context())
	c.Data["json"] = map[string]string{
		"request_id": state.RequestID,
		"trace_id":   state.TraceID,
		"session_id": state.SessionID,
	}
	c.ServeJSON()
}

// boomController panics with a value that must never reach the client.
type boomController struct {
	beego.Controller
}

func (c *boomController) Get() {
	panic("do-not-leak-secret-value")
}

// failController uses the respondError helper.
type failController struct {
	beego.Controller
}

func (c *failController) Get() {
	RespondError(&c.Controller, 400, "bad_request", "something wrong")
}

// stopController calls StopRun (beego aborts via panic(beego.ErrAbort)).
type stopController struct {
	beego.Controller
}

func (c *stopController) Get() {
	c.StopRun()
}

func newTestHandler(t *testing.T, logBuf *bytes.Buffer) http.Handler {
	t.Helper()
	beego.BConfig.RecoverFunc = nil
	if err := logging.InitWithWriter(logBuf); err != nil {
		t.Fatal(err)
	}

	cr := beego.NewControllerRegister()
	if err := cr.Add("/echo", &echoController{}, "get:Get"); err != nil {
		t.Fatal(err)
	}
	if err := cr.Add("/boom", &boomController{}, "get:Get"); err != nil {
		t.Fatal(err)
	}
	if err := cr.Add("/fail", &failController{}, "get:Get"); err != nil {
		t.Fatal(err)
	}
	if err := cr.Add("/stop", &stopController{}, "get:Get"); err != nil {
		t.Fatal(err)
	}
	if err := cr.Add("/v1/hello/:id", &echoController{}, "get:Get"); err != nil {
		t.Fatal(err)
	}
	if err := cr.InsertFilter("*", beego.BeforeExec, RoutePattern()); err != nil {
		t.Fatal(err)
	}
	return Trace(cr)
}

func accessLogLines(logBuf *bytes.Buffer) []map[string]interface{} {
	var out []map[string]interface{}
	for _, line := range strings.Split(strings.TrimSpace(logBuf.String()), "\n") {
		if line == "" {
			continue
		}
		m := map[string]interface{}{}
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			continue
		}
		if m["msg"] == "http_request" {
			out = append(out, m)
		}
	}
	return out
}

var generatedIDPattern = regexp.MustCompile(`^[0-9a-f]{32}$`)

func TestValidHeadersAreStoredAndEchoed(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	req := httptest.NewRequest("GET", "/echo", nil)
	req.Header.Set(tracing.HeaderSessionID, "session_123456")
	req.Header.Set(tracing.HeaderTraceID, "trace_12345678")
	req.Header.Set(tracing.HeaderRequestID, "request_12345678")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Fatalf("status = %d", rec.Code)
	}
	if got := rec.Header().Get(tracing.HeaderRequestID); got != "request_12345678" {
		t.Errorf("response X-Request-Id = %q", got)
	}
	if got := rec.Header().Get(tracing.HeaderTraceID); got != "trace_12345678" {
		t.Errorf("response X-Trace-Id = %q", got)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("bad body: %v", err)
	}
	if body["request_id"] != "request_12345678" || body["trace_id"] != "trace_12345678" || body["session_id"] != "session_123456" {
		t.Errorf("controller saw IDs %v", body)
	}

	lines := accessLogLines(&logBuf)
	if len(lines) != 1 {
		t.Fatalf("expected exactly one access log line, got %d", len(lines))
	}
	m := lines[0]
	if m["method"] != "GET" || m["path"] != "/echo" || m["status"] != float64(200) {
		t.Errorf("access line = %v", m)
	}
	if m["request_id"] != "request_12345678" || m["trace_id"] != "trace_12345678" || m["session_id"] != "session_123456" {
		t.Errorf("access line IDs = %v", m)
	}
	if m["service"] != logging.ServiceName {
		t.Errorf("access line service = %v", m["service"])
	}
}

func TestMissingHeadersAreGenerated(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/echo", nil))

	reqID := rec.Header().Get(tracing.HeaderRequestID)
	traceID := rec.Header().Get(tracing.HeaderTraceID)
	if !generatedIDPattern.MatchString(reqID) {
		t.Errorf("generated request ID invalid: %q", reqID)
	}
	if traceID != reqID {
		t.Errorf("trace ID should fall back to request ID: trace=%q request=%q", traceID, reqID)
	}

	var body map[string]string
	json.Unmarshal(rec.Body.Bytes(), &body)
	if body["request_id"] != reqID || body["trace_id"] != traceID || body["session_id"] != "" {
		t.Errorf("controller saw IDs %v (want request=%q trace=%q empty session)", body, reqID, traceID)
	}
}

func TestInvalidHeadersAreDiscarded(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	req := httptest.NewRequest("GET", "/echo", nil)
	// Injection attempts: each must be discarded, never logged raw.
	req.Header.Set(tracing.HeaderSessionID, "abc;DROP TABLE users")
	req.Header.Set(tracing.HeaderTraceID, "trace-id\nX-Evil: 1")
	req.Header.Set(tracing.HeaderRequestID, "short")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	reqID := rec.Header().Get(tracing.HeaderRequestID)
	if !generatedIDPattern.MatchString(reqID) {
		t.Errorf("invalid request ID not replaced by a generated one: %q", reqID)
	}
	if rec.Header().Get(tracing.HeaderTraceID) != reqID {
		t.Errorf("invalid trace ID should fall back to generated request ID, got %q", rec.Header().Get(tracing.HeaderTraceID))
	}
	for _, raw := range []string{"DROP TABLE users", "trace-id", "X-Evil", "short"} {
		if strings.Contains(logBuf.String(), raw) {
			t.Errorf("raw invalid header value %q leaked into logs", raw)
		}
	}
}

func TestPanicReturnsJSONErrorWithoutStack(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/boom", nil))

	if rec.Code != 500 {
		t.Fatalf("status = %d, want 500", rec.Code)
	}
	if strings.Contains(rec.Body.String(), "do-not-leak-secret-value") {
		t.Error("panic value leaked to client")
	}
	if strings.Contains(rec.Body.String(), "goroutine") {
		t.Error("stack trace leaked to client")
	}

	var body struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("error body is not the JSON envelope: %v (%s)", err, rec.Body.String())
	}
	if body.Error.Code != "internal_error" {
		t.Errorf("error code = %q", body.Error.Code)
	}
	if body.Error.RequestID == "" || body.Error.RequestID != rec.Header().Get(tracing.HeaderRequestID) {
		t.Errorf("error body request_id = %q, response header = %q", body.Error.RequestID, rec.Header().Get(tracing.HeaderRequestID))
	}

	// The panic itself must be logged with IDs and a stack, server-side only.
	if !strings.Contains(logBuf.String(), "panic recovered") {
		t.Error("panic not logged")
	}
	if !strings.Contains(logBuf.String(), "do-not-leak-secret-value") {
		t.Error("panic value not logged server-side")
	}
	if !strings.Contains(logBuf.String(), "stack") {
		t.Error("panic stack not logged server-side")
	}
	// Still exactly one access log line for the panicked request.
	if len(accessLogLines(&logBuf)) != 1 {
		t.Errorf("expected exactly one access log line after panic, got %d", len(accessLogLines(&logBuf)))
	}
}

func TestRespondErrorIncludesRequestID(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	req := httptest.NewRequest("GET", "/fail", nil)
	req.Header.Set(tracing.HeaderRequestID, "request_12345678")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	if rec.Code != 400 {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
	var body struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("bad error body: %v (%s)", err, rec.Body.String())
	}
	if body.Error.Code != "bad_request" || body.Error.Message != "something wrong" {
		t.Errorf("error body = %v", body.Error)
	}
	if body.Error.RequestID != "request_12345678" {
		t.Errorf("error body request_id = %q", body.Error.RequestID)
	}
	if rec.Header().Get(tracing.HeaderRequestID) != "request_12345678" {
		t.Errorf("response header X-Request-Id = %q", rec.Header().Get(tracing.HeaderRequestID))
	}
}

func TestEarlyNotFoundResponseGetsHeadersAndAccessLog(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/does-not-exist", nil))

	if rec.Code != 404 {
		t.Fatalf("status = %d, want 404", rec.Code)
	}
	if rec.Header().Get(tracing.HeaderRequestID) == "" {
		t.Error("404 response missing X-Request-Id")
	}
	if rec.Header().Get(tracing.HeaderTraceID) == "" {
		t.Error("404 response missing X-Trace-Id")
	}
	lines := accessLogLines(&logBuf)
	if len(lines) != 1 {
		t.Fatalf("expected exactly one access log line for 404, got %d", len(lines))
	}
	if lines[0]["status"] != float64(404) {
		t.Errorf("access line status = %v", lines[0]["status"])
	}
}

func TestBeegoAbortDoesNotProduce500(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/stop", nil))

	if rec.Code == 500 {
		t.Fatal("beego StopRun abort was treated as a panic")
	}
	if len(accessLogLines(&logBuf)) != 1 {
		t.Errorf("expected exactly one access log line, got %d", len(accessLogLines(&logBuf)))
	}
}

func TestAccessLogHasRoutePatternAndNoQueryValues(t *testing.T) {
	var logBuf bytes.Buffer
	h := newTestHandler(t, &logBuf)

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest("GET", "/v1/hello/123?password=supersecret", nil))

	lines := accessLogLines(&logBuf)
	if len(lines) != 1 {
		t.Fatalf("expected exactly one access log line, got %d", len(lines))
	}
	m := lines[0]
	if m["path"] != "/v1/hello/123" {
		t.Errorf("access path = %v (query values must not be logged)", m["path"])
	}
	if m["route"] != "/v1/hello/:id" {
		t.Errorf("access route = %v", m["route"])
	}
	if strings.Contains(logBuf.String(), "supersecret") {
		t.Error("query string value leaked into logs")
	}
	for _, forbidden := range []string{"authorization", "cookie", "password", "body"} {
		if _, ok := m[forbidden]; ok {
			t.Errorf("access log contains forbidden key %q", forbidden)
		}
	}
	if _, ok := m["duration_ms"]; !ok {
		t.Error("access line missing duration_ms")
	}
	if _, ok := m["bytes"]; !ok {
		t.Error("access line missing bytes")
	}
}

func TestWriteErrorShape(t *testing.T) {
	rec := httptest.NewRecorder()
	WriteError(rec, 422, "validation_failed", "bad input", "request_12345678")

	if rec.Code != 422 {
		t.Errorf("status = %d", rec.Code)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Errorf("content type = %q", ct)
	}
	var body struct {
		Error struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			RequestID string `json:"request_id"`
		} `json:"error"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("bad error body: %v", err)
	}
	if body.Error.Code != "validation_failed" || body.Error.Message != "bad input" || body.Error.RequestID != "request_12345678" {
		t.Errorf("error body = %v", body.Error)
	}
}
