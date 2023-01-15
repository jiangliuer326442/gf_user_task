package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

func main() {
	var (
		ctx = gctx.New()
	)
	cron := gcron.New()
	array := garray.New(true)
	cron.AddTimes(ctx, "@every 1s", 2, func(ctx context.Context) {
		array.Append(1)
	}, "cron1")
	cron.AddOnce(ctx, "@every 1s", func(ctx context.Context) {
		array.Append(1)
	}, "cron2")
	entries := cron.Entries()
	for k, v := range entries {
		fmt.Println(k, v.Name, v.Time)

	}
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(array.Len())

}
