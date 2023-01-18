package cmd

import (
	"context"
	userController "gf_user_task/internal/controller/user"
	myMiddleware "gf_user_task/internal/library/middleware"
	smcService "gf_user_task/internal/logic/user/smc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	rpcxServer "github.com/smallnest/rpcx/server"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s1 := g.Server("user")
			s1.Use(myMiddleware.HandleTracing)
			s1.Use(myMiddleware.HandleLang)
			s1.Use(myMiddleware.HandlerResponse)
			s1.Group("/user", func(group *ghttp.RouterGroup) {
				group.Bind(
					userController.Smc,
				)
			})

			if err = s1.Start(); err != nil {
				panic(err)
			}

			//s2 := krpc.Server.NewGrpcServer()
			//grpcSmcServoce.RegisterSmcServer(s2.Server, smcService.New())
			//s2.Start()

			s3 := rpcxServer.NewServer()
			s3.RegisterName("Smc", smcService.New(), "")
			go func() {
				s3.Serve("tcp", ":8772")
			}()

			g.Wait()
			return nil
		},
	}
)
