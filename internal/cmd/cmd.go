package cmd

import (
	"context"
	grpcSmcServoce "gf_user_task/generated/user/protobuf/smc"
	userController "gf_user_task/internal/controller/user"
	myMiddleware "gf_user_task/internal/library/middleware"
	smcService "gf_user_task/internal/logic/user/smc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/katyusha/krpc"
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

			s2 := krpc.Server.NewGrpcServer()
			grpcSmcServoce.RegisterSmcServer(s2.Server, smcService.New())
			s2.Start()

			g.Wait()
			return nil
		},
	}
)
