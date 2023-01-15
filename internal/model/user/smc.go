package user

import (
	"context"
	"database/sql"
	"gf_user_task/generated/user/model/entity"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

var (
	usersByNicknameCache = gcache.New(gconv.Int(g.Cfg().MustGet(context.Background(), "diy.user.usersByNicknameCacheSize", 10)))
)

func GetByNickname(ctx context.Context, nickname string) (userEntity *entity.Users, err error) {
	ctx, span := gtrace.NewSpan(ctx, "GetByNickname")
	defer span.End()
	var rawUserEntity *gvar.Var
	if rawUserEntity, err = usersByNicknameCache.Get(ctx, generateNicknamekey(nickname)); err == nil && rawUserEntity != nil {
		if err = gconv.Struct(rawUserEntity, &userEntity); err == nil && userEntity != nil {
			return userEntity, nil
		}
	}

	var gRedis = g.Redis("user")
	rawRedisResult, err2 := gRedis.Do(ctx, "GET", generateNicknamekey(nickname))
	if err2 == nil {
		if err = rawRedisResult.Struct(&userEntity); err == nil && userEntity != nil {
			return userEntity, nil
		}
	}

	var gModel = g.DB("user").Ctx(ctx).Model("users")
	if err = gModel.Where("name = ?", nickname).Scan(&userEntity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if _, err = gRedis.Do(ctx, "SET", generateNicknamekey(nickname), userEntity); err != nil {
		glog.Error(ctx, err)
	}
	if _, err = gRedis.Do(ctx, "EXPIRE", generateNicknamekey(nickname), 3600*8); err != nil {
		glog.Error(ctx, err)
	}

	if err = usersByNicknameCache.Set(ctx, generateNicknamekey(nickname), userEntity, time.Hour); err != nil {
		glog.Error(ctx, err)
	}

	return userEntity, nil
}

func AddUser(userEntity *entity.Users) (lastInsertId int64, err error) {
	var gModel = g.DB("user").Model("users")
	return gModel.Data(userEntity).InsertAndGetId()
}

func generateNicknamekey(nickname string) string {
	return "user_by_nickname_" + nickname
}
