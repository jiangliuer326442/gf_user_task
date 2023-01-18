package user

import (
	"context"
	CommonFields "gf_user_task/api/web/v1"
	userSmcApi "gf_user_task/api/web/v1/user/smc"
	grpcSmcServoce "gf_user_task/generated/user/protobuf/smc"
	"gf_user_task/internal/service/user"
	"github.com/gogf/gf/v2/frame/g"
	rpcxClient "github.com/smallnest/rpcx/client"
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
	errRaw := user.Jwt().Login(ctx, req, res)
	err = errRaw
	if err != nil {
		return
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

func (c *cSmc) Rpc2(ctx context.Context, req *userSmcApi.Rpc2Req) (res *userSmcApi.RegisterSimpleRes, err error) {

	// #1
	d, err := rpcxClient.NewPeer2PeerDiscovery("tcp@127.0.0.1:8772", "")
	if err != nil {
		return
	}
	// #2
	xclient := rpcxClient.NewXClient("Smc", rpcxClient.Failtry, rpcxClient.RandomSelect, d, rpcxClient.DefaultOption)
	defer xclient.Close()

	commonFileds := &CommonFields.Common{
		Lang: "en",
	}

	// #3
	args := &userSmcApi.LoginSimpleReq{
		Common:   *commonFileds,
		Nickname: "mustafa3264",
		Password: "123456",
	}

	// #4
	reply := &userSmcApi.RegisterSimpleRes{}

	// #5
	err = xclient.Call(ctx, "Login", args, reply)
	if err != nil {
		return
	}

	res = reply

	return
}
