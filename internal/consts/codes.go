package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	CodeNicknameRegistered = gcode.New(10001, `smcnickregistered`, nil)
	CodeNicknameNonExisted = gcode.New(10002, `smcnicknotexists`, nil)
	CodePasswordNotComple  = gcode.New(10003, `smcpwdnotcompable`, nil)
)
