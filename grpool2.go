package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"sync"
)

func main() {
	ctx := gctx.New()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		v := i
		grpool.Add(ctx, func(ctx context.Context) {
			fmt.Println(v)
			wg.Done()
		})
	}
	wg.Wait()
}
