package smc

import (
	CommonFields "gf_user_task/api/web/v1"
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterSimpleReq struct {
	g.Meta `path:"/register" tags:"Smc" method:"post" summary:"使用昵称和密码注册"`
	CommonFields.Common
	Nickname  string `v:"required|length:6,16#nickrequired|registernickerr" dc:"用户昵称" form:"nickname" example:"mustafa3264"`
	Password  string `v:"required|length:6,16#passwordrequired|registerpwderr" dc:"用户密码" form:"password" example:"12345678"`
	Password2 string `v:"same:Password#registerpwdrepeat" dc:"确认密码"`
}
type RegisterSimpleRes struct {
	Uid   int64  `dc:"用户uid" json:"uid" example:"224466"`
	Token string `dc:"用户token" json:"token" example:"abcdefghijkl"`
}

type LoginSimpleReq struct {
	g.Meta `path:"/login" tags:"Smc" method:"post" summary:"使用昵称和密码登陆"`
	CommonFields.Common
	Nickname string `v:"required#nickrequired" dc:"用户昵称" form:"nickname" example:"mustafa3264"`
	Password string `v:"required#passwordrequired" dc:"用户密码" form:"password" example:"12345678"`
}

type RpcReq struct {
	g.Meta `path:"/rpc" tags:"Rpc" method:"get" summary:"微服务调用登陆"`
}
