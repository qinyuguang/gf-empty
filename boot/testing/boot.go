package testing

import (
	"gf-empty/library/snowflake"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/genv"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
)

func init() {
	env := "dev"

	// set time zone
	gtime.SetTimeZone("Asia/Shanghai")

	// set config file
	setConfigFile(env)

	// set os env
	osSetenv()

	// set logger
	setLogger()

	// set snowflake
	setSnowflake(env)
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
	g.Log().SetStack(false)
	g.Log().SetStdoutPrint(false)
}

func setSnowflake(env string) {
	snowflake.Init(snowflake.NodeTypePrivate)
}
