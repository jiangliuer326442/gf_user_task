package main

import (
	_ "gf_user_task/internal/library/apollo"
	_ "gf_user_task/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"

	_ "gf_user_task/internal/logic/user"

	"gf_user_task/internal/cmd"
)

func main() {
	var ctx = gctx.New()
	var err error
	err = g.I18n().SetPath("resource/i18n")
	if err != nil {
		panic(err)
	}
	ServiceName, err := g.Cfg().Get(ctx, "grpcserver.name")
	if err != nil {
		panic(err)
	}
	JaegerUdpEndpoint, err := g.Cfg().Get(ctx, "JaegerUdpEndpoint")
	if err != nil {
		panic(err)
	}
	tp, err := jaeger.Init(gconv.String(ServiceName), gconv.String(JaegerUdpEndpoint))
	if err != nil {
		panic(err)
	}
	defer tp.Shutdown(ctx)

	cmd.Main.Run(ctx)
}
