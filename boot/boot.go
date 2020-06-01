package boot

import (
	"gf-empty/library/middleware"
	"gf-empty/library/rpc"
	rpc_debug "gf-empty/library/rpc/middleware/debug"
	rpc_recovery "gf-empty/library/rpc/middleware/recovery"
	rpc_requestid "gf-empty/library/rpc/middleware/requestid"
	"gf-empty/library/snowflake"
	"gf-empty/library/trace"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func init() {
	// args parser
	parser, err := gcmd.Parse(g.MapStrBool{
		"e,env": true,
	})
	if err != nil {
		g.Log().Error("parser err:", err)
		return
	}

	env := getEnv(parser)

	// set time zone
	gtime.SetTimeZone("Asia/Shanghai")

	// set config file
	setConfigFile(env)

	// set os env
	osSetenv()

	// set logger
	setLogger()

	// set middleware
	setMiddleware()

	// set RPC
	setRPC(env)

	// set snowflake
	setSnowflake(env)

	// set database
	setDatabase()
}

func getEnv(parser *gcmd.Parser) string {
	env := genv.Get("ENV")
	if env != "" {
		return env
	}

	env = parser.GetOpt("e", "dev")
	genv.Set("ENV", env)

	return env
}

func setConfigFile(env string) {
	filepath := "config/" + env + ".toml"

	// lookup config file for test
	for !gfile.Exists(filepath) {
		filepath = "../" + filepath
	}

	if gfile.Exists(filepath) {
		g.Cfg().SetFileName(env + ".toml")
	} else {
		g.Log().Fatal("config file not found")
	}
}

func osSetenv() {
	vars := g.Cfg().GetMapStrStr("osenv")
	if vars == nil {
		return
	}

	for name, value := range vars {
		if v := genv.Get(name); v != "" {
			genv.Set(name, v)
		} else {
			genv.Set(name, value)
		}
	}
}

func setLogger() {
	if level := genv.Get("LOG_LEVEL"); level != "" {
		g.Log().SetLevelStr(level)
		g.Cfg().Set("logger.level", level)
	}

	// set glog
	// g.Log().SetAsync(true)
}

func setMiddleware() {
	s := g.Server()

	s.BindMiddlewareDefault(middleware.Init)
	s.BindMiddlewareDefault(middleware.Log)
}

func setRPC(env string) {
	_, err := trace.InitGlobalTracer()
	if err != nil {
		g.Log().Fatal("InitGlobalTracer err", err)
	}

	rpc.Server().Options(
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(
				grpc_recovery.WithRecoveryHandler(rpc_recovery.Logger),
			),
			rpc_requestid.UnaryServerInterceptor(
				rpc_requestid.WithScope(genv.Get("APP_NAME")),
			),
			grpc_validator.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
		),
	)

	if g.Cfg().GetString("logger.level") == "debug" || g.Cfg().GetString("logger.level") == "all" {
		rpc.Server().Options(
			grpc.ChainUnaryInterceptor(
				rpc_debug.UnaryServerInterceptor(),
			),
		)
	}
}

func setSnowflake(env string) {
	snowflake.Init(snowflake.NodeTypePrivate)
}

func setDatabase() {
	// db logger
	g.DB().SetLogger(g.Log())
}
