package smc

import (
	"context"
	"gf_user_task/generated/user/model/entity"
	grpcSmcServoce "gf_user_task/generated/user/protobuf/smc"
	"gf_user_task/internal/consts"
	"gf_user_task/internal/model/user"
	userService "gf_user_task/internal/service/user"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"strconv"
)

type (
	NickPwdUser struct {
		grpcSmcServoce.UnimplementedSmcServer
	}
)

func init() {
	userService.RegisterJwt(New())
}

func New() *NickPwdUser {
	return &NickPwdUser{}
}

func (u *NickPwdUser) Register(gcontext context.Context, q *grpcSmcServoce.RegisterSimpleReq) (r *grpcSmcServoce.RegisterSimpleRes, err error) {
	gcontext, span := gtrace.NewSpan(gcontext, "Register")
	defer span.End()
	userEntity, err3 := user.GetByNickname(gcontext, q.Nickname)
	if err3 != nil {
		return nil, err3
	}
	if userEntity != nil {
		return nil, gerror.NewCode(consts.CodeNicknameRegistered, g.I18n().Translate(gcontext, consts.CodeNicknameRegistered.Message()))
	}

	pwdEncrypted, err4 := generatePwd(gcontext, q.Password)
	if err4 != nil {
		return nil, err4
	}

	userEntity = &entity.Users{
		Name:     q.Nickname,
		Password: pwdEncrypted,
	}
	uid, err1 := user.AddUser(userEntity)
	if err1 != nil {
		return nil, err1
	}

	token, err3 := generateToken(gcontext, uint(uid))
	if err3 != nil {
		return nil, err3
	}

	return &grpcSmcServoce.RegisterSimpleRes{
		Uid:   uid,
		Token: token,
	}, nil
}

func (u *NickPwdUser) Login(gcontext context.Context, q *grpcSmcServoce.LoginSimpleReq) (r *grpcSmcServoce.RegisterSimpleRes, err error) {
	gcontext, span := gtrace.NewSpan(gcontext, "Login")
	defer span.End()
	userEntity, err3 := user.GetByNickname(gcontext, q.Nickname)
	if err3 != nil {
		return nil, err3
	}
	if userEntity == nil {
		return nil, gerror.NewCode(consts.CodeNicknameNonExisted, g.I18n().Translate(gcontext, consts.CodeNicknameNonExisted.Message()))
	}

	pwdEncrypted, err4 := generatePwd(gcontext, q.Password)
	glog.Debug(gcontext, "passwordmd5", q.Password, pwdEncrypted)
	if err4 != nil {
		return nil, err4
	}

	if pwdEncrypted != userEntity.Password {
		return nil, gerror.NewCode(consts.CodePasswordNotComple, g.I18n().Translate(gcontext, consts.CodePasswordNotComple.Message()))
	}
	uid := userEntity.Id

	token, err3 := generateToken(gcontext, uid)
	if err3 != nil {
		return nil, err3
	}

	return &grpcSmcServoce.RegisterSimpleRes{
		Uid:   int64(uid),
		Token: token,
	}, nil
}

func generatePwd(ctx context.Context, pwd string) (string, error) {
	saltPwd, err4 := g.Cfg().Get(ctx, "jwtPwd")
	if err4 != nil {
		return "", err4
	}
	pwdEncrypted, err5 := gmd5.Encrypt(gconv.String(saltPwd) + pwd)
	if err5 != nil {
		return "", err5
	}
	return pwdEncrypted, nil
}

func generateToken(ctx context.Context, uid uint) (string, error) {
	salt, err2 := g.Cfg().Get(ctx, "jwtToken")
	if err2 != nil {
		return "", err2
	}
	token, err3 := gmd5.Encrypt(gconv.String(salt) + strconv.FormatUint(uint64(uid), 10))
	if err3 != nil {
		return "", err3
	}
	glog.Debug(ctx, "md5字符串"+gconv.String(salt)+strconv.FormatUint(uint64(uid), 10))
	return token, nil
}
