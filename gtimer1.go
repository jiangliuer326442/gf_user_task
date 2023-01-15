package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"time"
)

func main() {
	var (
		ctx      = gctx.New()
		timeout  = time.Second
		interval = time.Second
	)
	gtimer.SetTimeout(ctx, timeout, func(ctx context.Context) {
		fmt.Println("SetTimeout:", gtime.Now())
	})
	gtimer.SetInterval(ctx, interval, func(ctx context.Context) {
		fmt.Println("SetInterval:", gtime.Now())
	})
	select {}
}
