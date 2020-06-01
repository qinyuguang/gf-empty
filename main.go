package main

import (
	"context"
	"sync"

	_ "gf-empty/boot"
	_ "gf-empty/router"

	"gf-empty/app/jobs"
	"gf-empty/library/rpc"

	"github.com/gogf/gf/frame/g"
)

func main() {
	// HTTP server
	go g.Server().Run()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// background worker
	jobs.RunAll(ctx, &wg)

	// GRPC server
	if err := rpc.Server().Run(ctx, &wg); err != nil {
		g.Log().Fatal("err:", err)
	} else {
		cancel()
	}

	wg.Wait()

}
