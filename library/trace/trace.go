package trace

import (
	"io"

	"github.com/gogf/gf/os/genv"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitGlobalTracer() (io.Closer, error) {
	var cfg *jaegercfg.Configuration

	cfg, err := jaegercfg.FromEnv()
	if err != nil {
		return nil, err
	}

	return cfg.InitGlobalTracer(
		genv.Get("APP_NAME"),
	)
}
