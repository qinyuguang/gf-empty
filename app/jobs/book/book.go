package book

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/grpool"
)

func Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	pool := grpool.New(10)

	for {
		select {
		case <-ctx.Done():
			g.Log().Info("book worker terming")
			for pool.Jobs() > 0 {
				g.Log().Debug("book worker still running. jobs count:", pool.Jobs())
				time.Sleep(500 * time.Millisecond)
			}
			g.Log().Info("book worker shutdown")
			return
		default:
			pool.Add(func() {
				// do something
				g.Log().Debug("book worker running")
			})
			time.Sleep(3 * time.Second)
		}
	}
}
