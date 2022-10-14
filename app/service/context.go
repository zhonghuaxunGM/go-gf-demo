package service

import (
	"context"
	"go-gf-demo/app/model"

	"github.com/gogf/gf/net/ghttp"
)

var Context = new(serviceContext)

type serviceContext struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *serviceContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *serviceContext) Get(ctx context.Context) *model.Context {
	val := ctx.Value(model.ContextKey)
	if val == nil {
		return nil
	}
	if localCtx, ok := val.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *serviceContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
