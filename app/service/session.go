package service

import (
	"context"
	"fmt"
	"go-gf-demo/app/model"
)

// Session管理服务
var Session = new(serviceSession)

type serviceSession struct{}

const (
	// 用户信息存放在Session中的Key
	sessionKeyUser = "SessionKeyUser"
)

// 判断用户登录是否成功，其session是否被分配，没有被分配则无法在上层封装 model.User
// 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *serviceSession) GetUser(ctx context.Context) *model.User {
	// 获得上下文变量 key=model.ContextKey, val=customCtx(*model.Context)
	customCtx := Context.Get(ctx)
	fmt.Println(customCtx)
	if customCtx != nil {
		// customCtx.Session=r.Session
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var user *model.User
			v.Struct(&user)
			return user
		}
	}
	return nil
}

// 设置用户Session.
func (s *serviceSession) SetUser(ctx context.Context, user *model.User) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}
