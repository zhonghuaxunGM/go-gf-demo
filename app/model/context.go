package model

import "github.com/gogf/gf/net/ghttp"

const (
	// 上下文变量存储键名
	ContextKey = "ContextKey"
)

// 请求上下文结构
type Context struct {
	Session *ghttp.Session
	User    *ContextUser
}

// 请求上下文中的用户信息
type ContextUser struct {
	Id       uint
	Passport string
	Nickname string
}
