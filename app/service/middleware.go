package service

import (
	"go-gf-demo/app/model"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

type serivceMiddleware struct {
}

var Middleware = new(serivceMiddleware)

// 自定义上下文对象
func (e *serivceMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	// context.WithValue  key=model.ContextKey, val=customCtx
	Context.Init(r, customCtx)
	// 通过 Session.GetUser(r.Context()) 封装 customCtx
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Passport: user.Passport,
			Nickname: user.Nickname,
		}
	}
	r.Middleware.Next()
}

// 鉴权中间件，只有登录成功（model.user封装完成）之后才能通过
func (e *serivceMiddleware) Auth(r *ghttp.Request) {
	// 通过SetParam和GetParam来设置和获取自定义的变量，该变量生命周期仅限于当前请求流程
	r.SetParam("site", "https://golang.org")
	if User.IsSignedIn(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}
