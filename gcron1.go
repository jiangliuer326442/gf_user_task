package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

func main() {
	var (
		err error
		ctx = gctx.New()
	)
	_, err = gcron.Add(ctx, "* * * * * *", func(ctx context.Context) {
		g.Log().Print(ctx, "Every second")
	}, "MySecondCronJob")
	if err != nil {
		panic(err)
	}

	_, err = gcron.Add(ctx, "0 30 * * * *", func(ctx context.Context) {
		g.Log().Print(ctx, "Every hour on the half hour")
	})
	if err != nil {
		panic(err)
	}

	_, err = gcron.Add(ctx, "@hourly", func(ctx context.Context) {
		g.Log().Print(ctx, "Every hour")
	})
	if err != nil {
		panic(err)
	}

	_, err = gcron.Add(ctx, "@every 1h30m", func(ctx context.Context) {
		g.Log().Print(ctx, "Every hour thirty")
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(gcron.Entries())
	//g.Dump(gcron.Entries())

	time.Sleep(3 * time.Second)

	g.Log().Print(ctx, `stop cronjob "MySecondCronJob"`)
	gcron.Stop("MySecondCronJob")

	time.Sleep(15 * time.Second)

	g.Log().Print(ctx, `start cronjob "MySecondCronJob"`)
	gcron.Start("MySecondCronJob")

	time.Sleep(20 * time.Second)
}
