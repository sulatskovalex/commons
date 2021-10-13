package cmctx

import (
	"context"
	"github.com/sulatskovalex/commons/errs"
)

const (
	userIdKey  = "uid"
	userTypKey = "utype"
)

type UserInfo struct {
	UserId   int64
	UserType int64
}

func GetUserInfo(ctx context.Context) *UserInfo {
	userId, _ := ctx.Value(userIdKey).(int64)
	userType, _ := ctx.Value(userTypKey).(int64)
	return &UserInfo{UserId: userId, UserType: userType}
}
func RequireUserInfo(ctx context.Context) (*UserInfo, error) {
	userId, _ := ctx.Value(userIdKey).(int64)
	if userId == 0 {
		return nil, errs.AccessDeniedErr
	}
	userType, _ := ctx.Value(userTypKey).(int64)
	return &UserInfo{UserId: userId, UserType: userType}, nil
}
func RequireUserType(ctx context.Context, uType int64) (int64, error) {
	userType, _ := ctx.Value(userTypKey).(int64)
	if userType <= uType {
		return 0, errs.AccessDeniedErr
	}
	userId, _ := ctx.Value(userIdKey).(int64)
	if userId == 0 {
		return 0, errs.AccessDeniedErr
	}
	return userId, nil
}
func UserInfoCtx(ctx context.Context, info *UserInfo) context.Context {
	return WithUserInfo(ctx, info)
}
func WithUserInfo(ctx context.Context, info *UserInfo) context.Context {
	return context.WithValue(context.WithValue(ctx, userIdKey, info.UserId), userTypKey, info.UserType)
}
func WithUserInfos(ctx context.Context, userId, userType int64) context.Context {
	return context.WithValue(context.WithValue(ctx, userIdKey, userId), userTypKey, userType)
}
