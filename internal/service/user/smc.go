package user

import (
	"context"
	grpcSmcServoce "gf_user_task/generated/user/protobuf/smc"
)

type (
	IJwt interface {
		Register(context.Context, *grpcSmcServoce.RegisterSimpleReq) (*grpcSmcServoce.RegisterSimpleRes, error)
		Login(context.Context, *grpcSmcServoce.LoginSimpleReq) (*grpcSmcServoce.RegisterSimpleRes, error)
	}
)

var (
	localJwt IJwt
)

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}
