package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"

	"customer_management_service/tracing"
)

func logLines(buf *bytes.Buffer) []map[string]interface{} {
	var out []map[string]interface{}
	for _, line := range strings.Split(strings.TrimSpace(buf.String()), "\n") {
		if line == "" {
			continue
		}
		m := map[string]interface{}{}
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			continue
		}
		out = append(out, m)
	}
	return out
}

func TestForIncludesServiceAndIDs(t *testing.T) {
	var buf bytes.Buffer
	if err := InitWithWriter(&buf); err != nil {
		t.Fatal(err)
	}

	state := &tracing.State{SessionID: "s-12345678", TraceID: "t-12345678", RequestID: "r-12345678"}
	ctx := tracing.NewContext(context.Background(), state)
	For(ctx).Info("hello", "extra", "value")

	lines := logLines(&buf)
	if len(lines) != 1 {
		t.Fatalf("expected 1 log line, got %d", len(lines))
	}
	m := lines[0]
	if m["service"] != ServiceName {
		t.Errorf("service = %v, want %v", m["service"], ServiceName)
	}
	if m["session_id"] != "s-12345678" || m["trace_id"] != "t-12345678" || m["request_id"] != "r-12345678" {
		t.Errorf("IDs missing from log line: %v", m)
	}
	if m["msg"] != "hello" || m["extra"] != "value" {
		t.Errorf("unexpected log line: %v", m)
	}
}

func TestForWithoutTracingState(t *testing.T) {
	var buf bytes.Buffer
	if err := InitWithWriter(&buf); err != nil {
		t.Fatal(err)
	}

	For(context.Background()).Info("startup")
	lines := logLines(&buf)
	if len(lines) != 1 {
		t.Fatalf("expected 1 log line, got %d", len(lines))
	}
	m := lines[0]
	if m["service"] != ServiceName {
		t.Errorf("service = %v", m["service"])
	}
	if _, ok := m["session_id"]; !ok {
		t.Error("session_id key missing")
	}
	if _, ok := m["trace_id"]; !ok {
		t.Error("trace_id key missing")
	}
	if _, ok := m["request_id"]; !ok {
		t.Error("request_id key missing")
	}
}

func TestJSONOutputIsOneObjectPerLine(t *testing.T) {
	var buf bytes.Buffer
	if err := InitWithWriter(&buf); err != nil {
		t.Fatal(err)
	}

	Logger().Info("first")
	For(context.Background()).Info("second", "k", "v")

	lines := logLines(&buf)
	if len(lines) != 2 {
		t.Fatalf("expected 2 JSON lines, got %d", len(lines))
	}
	if lines[0]["msg"] != "first" || lines[1]["msg"] != "second" {
		t.Errorf("unexpected lines: %v", lines)
	}
}
