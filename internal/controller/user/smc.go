package user

import (
	"context"
	userSmcApi "gf_user_task/api/web/v1/user/smc"
	grpcSmcServoce "gf_user_task/generated/user/protobuf/smc"
	"gf_user_task/internal/service/user"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Smc = cSmc{}
)

type cSmc struct{}

func (c *cSmc) Register(ctx context.Context, req *userSmcApi.RegisterSimpleReq) (res *userSmcApi.RegisterSimpleRes, err error) {
	resRaw, errRaw := user.Jwt().Register(ctx, &grpcSmcServoce.RegisterSimpleReq{
		Lang:     req.Lang,
		Nickname: req.Nickname,
		Password: req.Password,
	})
	err = errRaw
	if err != nil {
		return
	}
	res = &userSmcApi.RegisterSimpleRes{
		Uid:   resRaw.Uid,
		Token: resRaw.Token,
	}

	return
}

func (c *cSmc) Login(ctx context.Context, req *userSmcApi.LoginSimpleReq) (res *userSmcApi.RegisterSimpleRes, err error) {
	resRaw, errRaw := user.Jwt().Login(ctx, &grpcSmcServoce.LoginSimpleReq{
		Lang:     req.Lang,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	err = errRaw
	if err != nil {
		return
	}
	res = &userSmcApi.RegisterSimpleRes{
		Uid:   resRaw.Uid,
		Token: resRaw.Token,
	}

	return
}

func (c *cSmc) Rpc(ctx context.Context, req *userSmcApi.RpcReq) (res *userSmcApi.RegisterSimpleRes, err error) {
	client, err := user.NewClient()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	resRaw, err := client.Smc().Login(ctx, &grpcSmcServoce.LoginSimpleReq{
		Lang:     "en",
		Nickname: "mustafa3264",
		Password: "123456",
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	res = &userSmcApi.RegisterSimpleRes{
		Uid:   resRaw.Uid,
		Token: resRaw.Token,
	}

	return
}
