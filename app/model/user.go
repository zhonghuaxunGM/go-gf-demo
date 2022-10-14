package model

import "github.com/gogf/gf/os/gtime"

type User struct {
	Id       uint        `orm:"id,primary" json:"id"`        // 用户ID
	Passport string      `orm:"passport"   json:"passport"`  // 用户账号
	Password string      `orm:"password"   json:"password"`  // 用户密码
	Nickname string      `orm:"nickname"   json:"nickname"`  // 用户昵称
	CreateAt *gtime.Time `orm:"create_at"  json:"create_at"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"update_at"` // 更新时间
}

type ApiUserSignInReq struct {
	Passport string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}
