package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"os"
	"time"
)

func main() {
	fmt.Printf("%d: I am child? %v\n", gproc.Pid(), gproc.IsChild())
	if gproc.IsChild() {
		ctx := context.Background()
		g.Log().Debug(ctx, `this is sub process`)
		gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
			err := gproc.Send(gproc.PPid(), []byte(gtime.Datetime()))
			if err != nil {
				return
			}
		})
		select {}
	} else {
		ctx := gctx.New()
		g.Log().Debug(ctx, `this is main process`)
		m := gproc.NewManager()
		g.Log().Debug(ctx, `env`, os.Environ())
		p := m.NewProcess(os.Args[0], os.Args, os.Environ())
		p.Start(ctx)
		for {
			msg := gproc.Receive()
			fmt.Printf("receive from %d, data: %s\n", msg.SenderPid, string(msg.Data))
		}
	}
}
