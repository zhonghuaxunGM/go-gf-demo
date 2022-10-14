package main

import (
	"fmt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcfg"
)

func cfgDemo() {
	// 读取config文件信息
	c := g.Cfg()
	fmt.Println(c.SetFileName("configdemo.toml"))
	Intojson(c.Get("db.def.1"))
	// 奇怪的事情发生了???
	fmt.Println(c.Get("title"))
	// c.AddPath("demo") //仅仅是多个搜索目录，检索到第一个文件便会停止
	// fmt.Println(c.Get("db2"))

	// 奇怪的事情发生了???
	fmt.Println(gcfg.GetContent("configdemo.toml"))
	gcfg.SetContent("title = cnxpaas", "configdemo.toml")
	gcfg.SetContent("db.def.0.user = cnxpaas", "configdemo.toml")
}
