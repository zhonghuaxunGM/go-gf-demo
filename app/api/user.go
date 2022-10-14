package api

import (
	"go-gf-demo/app/model"
	"go-gf-demo/app/service"
	"go-gf-demo/library/response"

	"github.com/gogf/gf/net/ghttp"
)

var User = new(apiUser)

type apiUser struct{}

func (a *apiUser) SignIn(r *ghttp.Request) {
	var data *model.ApiUserSignInReq
	// 任意的解析出model.ApiUserSignInReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.SignIn(r.Context(), data.Passport, data.Password); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *apiUser) Profile(r *ghttp.Request) {
	service.User.GetProfile(r.Context())
	// JsonExit(r, 0, "", service.User.GetProfile(r.Context()))
}
