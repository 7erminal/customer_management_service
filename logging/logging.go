// Package logging provides structured JSON logging via log/slog.
//
// Every log line carries service, session_id, trace_id and request_id.
// logging.For(ctx) returns a logger pre-loaded with the tracing
// identifiers stored in the request context by middleware.Trace.
package logging

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"customer_management_service/tracing"

	beego "github.com/beego/beego/v2/server/web"
)

// ServiceName identifies this service in every log line.
const ServiceName = "customer_management_service"

var (
	baseLogger *slog.Logger
	levelVar   slog.LevelVar
	logFile    *os.File
)

// Init configures the application logger. Output goes to stdout and to
// the log file (LOG_FILE env overrides the default location). The level
// comes from LOG_LEVEL (debug|info|warn|error), falling back to the
// `loglevel` key in conf/app.conf, then to info.
func Init() error {
	levelVar.Set(parseLevel(os.Getenv("LOG_LEVEL"), defaultConfLevel()))

	path := os.Getenv("LOG_FILE")
	if path == "" {
		path = "../logs/customer_management_application.log"
	}

	writers := []io.Writer{os.Stdout}
	if f, err := openLogFile(path); err == nil {
		logFile = f
		writers = append(writers, f)
	} else {
		// File logging is best effort; stdout still carries everything.
		fmt.Fprintf(os.Stderr, "logging: cannot open log file %q: %v\n", path, err)
	}

	_ = InitWithWriter(io.MultiWriter(writers...))
	return nil
}

// InitWithWriter configures the application logger writing to w. It is
// the programmatic form of Init and is used by tests.
func InitWithWriter(w io.Writer) error {
	baseLogger = slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: &levelVar}))
	slog.SetDefault(baseLogger)
	return nil
}

// Close releases the log file, if one was opened.
func Close() {
	if logFile != nil {
		_ = logFile.Close()
	}
}

// For returns a logger carrying the tracing identifiers from ctx. When
// ctx has no tracing state the IDs are logged as empty strings so every
// line keeps the same keys.
func For(ctx context.Context) *slog.Logger {
	if baseLogger == nil {
		_ = Init()
	}
	sessionID, traceID, requestID := "", "", ""
	if state, ok := tracing.FromContext(ctx); ok && state != nil {
		sessionID, traceID, requestID = state.SessionID, state.TraceID, state.RequestID
	}
	return baseLogger.With(
		slog.String("service", ServiceName),
		slog.String("session_id", sessionID),
		slog.String("trace_id", traceID),
		slog.String("request_id", requestID),
	)
}

// Logger returns the application logger without tracing identifiers
// (empty values) for use outside the request path, e.g. at startup.
func Logger() *slog.Logger {
	return For(context.Background())
}

func defaultConfLevel() string {
	v, err := beego.AppConfig.String("loglevel")
	if err != nil {
		return ""
	}
	return v
}

func parseLevel(values ...string) slog.Level {
	for _, v := range values {
		v = strings.ToLower(strings.TrimSpace(v))
		if v == "" {
			continue
		}
		switch v {
		case "debug":
			return slog.LevelDebug
		case "info":
			return slog.LevelInfo
		case "warn", "warning":
			return slog.LevelWarn
		case "error":
			return slog.LevelError
		}
	}
	return slog.LevelInfo
}

func openLogFile(path string) (*os.File, error) {
	if dir := dirOf(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return nil, err
		}
	}
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
}

func dirOf(path string) string {
	i := strings.LastIndexByte(path, '/')
	if i <= 0 {
		return ""
	}
	return path[:i]
}
