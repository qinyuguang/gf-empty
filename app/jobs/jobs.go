package jobs

import (
	"context"
	"sync"

	"gf-empty/app/jobs/book"
)

func RunAll(ctx context.Context, wg *sync.WaitGroup) {
	// 查询发送结果
	wg.Add(1)
	go book.Run(ctx, wg)
}
