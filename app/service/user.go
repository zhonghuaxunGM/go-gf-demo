package service

import (
	"context"
	"go-gf-demo/app/model"
)

// 中间件管理服务
var User = new(serviceUser)

type serviceUser struct{}

// 判断用户是否已经登录
// 判断用户 model.User是否正确封装完成
func (s *serviceUser) IsSignedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *serviceUser) SignIn(ctx context.Context, passport, password string) error {
	// user 认证通过...
	user := model.User{
		Id:       1,
		Passport: "psp",
		Password: "pwd",
		Nickname: "nick",
	}
	if err := Session.SetUser(ctx, &user); err != nil {
		return nil
	}
	Context.SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Passport: user.Passport,
		Nickname: user.Nickname,
	})
	return nil
}

// 获得用户信息详情
func (s *serviceUser) GetProfile(ctx context.Context) *model.User {
	return Session.GetUser(ctx)
}
