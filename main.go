package main

import (
	"customer_management_service/internal/trace"
	_ "customer_management_service/routers"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"

	beecontext "github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/server/web/filter/cors"
)

func configureTracing() {
	level := slog.LevelInfo
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "debug":
		level = slog.LevelDebug
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})))
	// Legacy Beego logs include sensitive request data; keep them disabled until
	// each call has been migrated to a context-aware, allowlisted event.
	logs.SetLevel(-1)
	// The standard logger is used by framework internals with untrusted text.
	log.SetOutput(io.Discard)
	orm.DebugLog.SetOutput(io.Discard)
	beego.BConfig.Log.AccessLogs = false
	// Beego uses ErrAbort as control flow; let the outer wrapper recover other panics.
	beego.BConfig.RecoverFunc = func(_ *beecontext.Context, _ *beego.Config) {
		if p := recover(); p != nil && p != beego.ErrAbort {
			panic(p)
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, func(c *beecontext.Context) {
		ids := trace.FromContext(c.Request.Context())
		if ids.RequestID == "" {
			ids = trace.IDsFromHeaders(c.Request.Header)
		}
		c.Request = c.Request.WithContext(trace.WithIDs(c.Request.Context(), ids))
		c.ResponseWriter.Header().Set(trace.HeaderRequestID, ids.RequestID)
		c.ResponseWriter.Header().Set(trace.HeaderTraceID, ids.TraceID)
	})
	beego.InsertFilter("*", beego.BeforeExec, func(c *beecontext.Context) {
		if st := trace.StateFromContext(c.Request.Context()); st != nil {
			if route, ok := c.Input.GetData("RouterPattern").(string); ok {
				st.Route = route
			}
		}
	})
}

func main() {
	configureTracing()
	sqlConn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		trace.Logger(nil, trace.Service).Error("database configuration unavailable")
	}

	if err := orm.RegisterDataBase("default", "mysql", sqlConn); err != nil {
		trace.Logger(nil, trace.Service).Error("database registration failed")
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174", "http://localhost:3000", "http://localhost:8000", "http://152.67.134.169", "http://13.40.60.131:8001", "http://185.249.227.127:9001", "http://185.249.227.127:9002", "http://167.86.115.44:8002", "http://5.252.55.191", "makufoodsltd.net", "https://makufoodsltd.net", "https://www.makufoodsltd.com", "https://makufoodsltd.com", "makufoodsltd.com", "https://admin.bridgeafrica.group", "https://mestechgh.com", "https://admin.mestechgh.com", "https://client.mestechgh.com", "https://authentication.mestechgh.com"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "X-Session-Id", "X-Trace-Id", "X-Request-Id"},
		ExposeHeaders:    []string{"Content-Length", "X-Request-Id", "X-Trace-Id"},
		AllowCredentials: true,
	}))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.Debug = false
	beego.RunWithMiddleWares("", trace.Middleware)
}
