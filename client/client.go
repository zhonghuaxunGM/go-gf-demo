package main

import (
	"fmt"
	"go-gf-demo/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func main() {
	// get请求
	fmt.Println("============get请求===========")
	if r, err := g.Client().Get("http://127.0.0.1:8300/register"); err != nil {
		panic(err)
	} else {
		defer r.Close()
		fmt.Println(r.ReadAllString())
	}
	// 下载远程文件
	// 文件过大，客户端发多个GET请求，每一次通过Header来请求分批的文件范围长度
	fmt.Println("============下载远程文件===========")
	if r, err := g.Client().Get("https://goframe.org/cover.png"); err != nil {
		panic(err)
	} else {
		defer r.Close()
		gfile.PutBytes("./cover.png", r.ReadAll())
	}
	// post请求
	fmt.Println("============post请求&符号===========")
	if r, err := g.Client().Post("http://127.0.0.1:8300/register", "username=user1&password1=pwd1&password2=passwd1"); err != nil {
		panic(err)
	} else {
		defer r.Close()
		fmt.Println(r.ReadAllString())
	}
	fmt.Println("============post请求map类型===========")
	if r, err := g.Client().Post("http://127.0.0.1:8300/register",
		g.Map{
			"username":  "user2",
			"password1": "pwd2",
			"password2": "passwd2",
		},
	); err != nil {
		panic(err)
	} else {
		defer r.Close()
		fmt.Println(r.ReadAllString())
	}
	fmt.Println("============post请求JSON数据===========")
	if r, err := g.Client().Post("http://127.0.0.1:8300/register",
		`{"username":"john","password1":"123456","password2":"123456"}`,
	); err != nil {
		panic(err)
	} else {
		defer r.Close()
		r.RawDump()
		fmt.Println(r.ReadAllString())
	}
	// 返回content为string类型
	fmt.Println("===========返回content为string类型============")
	g.Client().SetCookieMap(g.MapStrStr{
		"name":  "V",
		"score": "sir",
	})
	g.Client().SetHeaderMap(g.MapStrStr{
		"Span_id": "0.0.1",
	})
	content := g.Client().PostContent(
		"http://127.0.0.1:8300/register",
		`{"username":"john","password1":"123456","password2":"123456"}`,
	)
	fmt.Println(content)
	fmt.Println("===========GetVar Scan============")
	var r response.JsonResponse
	g.Client().GetVar(
		"http://127.0.0.1:8300/register",
		`{"username":"john","password1":"123456","password2":"123456"}`,
	).Scan(&r)
	fmt.Println(r)

	fmt.Println("===========upload files============")
	path := "./tmp/temp.txt"
	res, err := g.Client().Post("http://127.0.0.1:8400/upload", "upload-file=@file:"+path)
	if err != nil {
		glog.Error(err)
	} else {
		defer res.Close()
		fmt.Println(string(res.ReadAll()))
	}
}
