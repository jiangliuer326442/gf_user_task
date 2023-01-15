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
		ctx = gctx.New()
	)
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		fmt.Println("exit:", gtime.Now())
		gtimer.Exit()
	})
	select {}
}
