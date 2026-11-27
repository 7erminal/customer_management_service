package main

import (
	"os"

	"customer_management_service/logging"
	"customer_management_service/middleware"
	_ "customer_management_service/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"

	"github.com/beego/beego/v2/server/web/filter/cors"
)

func main() {
	// Structured JSON logging (log/slog). Every application log line
	// carries service, session_id, trace_id and request_id. Level comes
	// from LOG_LEVEL env, then the `loglevel` key in conf/app.conf,
	// then defaults to info.
	if err := logging.Init(); err != nil {
		logs.Error("failed to initialise logging: %s", err)
	}
	defer logging.Close()

	sqlConn, err := beego.AppConfig.String("sqlconn")
	if err != nil {
		logging.Logger().Error("failed to read sqlconn", "error", err.Error())
	}

	orm.RegisterDataBase("default", "mysql", sqlConn)
	logs.SetLogger(logs.AdapterFile, `{"filename":"../logs/customer_management_application.log"}`)

	// Suppress beego's built-in dev access-log line (logs.Debug): we
	// emit exactly one structured access-log line per request via
	// middleware.Trace instead.
	logs.SetLevel(logs.LevelInformational)

	// Disable beego's built-in panic recovery: in dev mode it renders
	// stack traces to the client. middleware.Trace provides recovery
	// with a JSON error envelope and no stack traces in responses.
	beego.BConfig.RecoverFunc = nil

	// ORM SQL debugging prints full SQL including parameter values
	// (PII). Off by default; opt in with LOG_ORM_DEBUG=true.
	orm.Debug = os.Getenv("LOG_ORM_DEBUG") == "true"

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174", "http://localhost:3000", "http://localhost:8000", "http://152.67.134.169", "http://13.40.60.131:8001", "http://185.249.227.127:9001", "http://185.249.227.127:9002", "http://167.86.115.44:8002", "http://5.252.55.191", "makufoodsltd.net", "https://makufoodsltd.net", "https://www.makufoodsltd.com", "https://makufoodsltd.com", "makufoodsltd.com", "https://admin.bridgeafrica.group", "https://mestechgh.com", "https://admin.mestechgh.com", "https://client.mestechgh.com", "https://authentication.mestechgh.com"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Record the matched route pattern for the access-log line.
	beego.InsertFilter("*", beego.BeforeExec, middleware.RoutePattern())

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	logging.Logger().Info("Starting service", "port", beego.BConfig.Listen.HTTPPort, "runmode", beego.BConfig.RunMode)

	// middleware.Trace wraps the whole beego handler: it validates and
	// propagates the gateway tracing headers, sets X-Request-Id /
	// X-Trace-Id on responses, recovers panics and writes the single
	// access-log line (including for panics and early responses).
	beego.RunWithMiddleWares("", middleware.Trace)
}
