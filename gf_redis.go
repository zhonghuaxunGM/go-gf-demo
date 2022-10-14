package main

import (
	"fmt"

	"github.com/gogf/gcache-adapter/adapter"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
)

func redisDemo() {
	c := gcache.New()

	c.Set("k1", "v1", 0)

	v, err := c.Get("k1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	s, err := c.Size()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	va1, err := c.Remove("k1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(va1)

	b, _ := c.Contains("k1")
	fmt.Println(b)
	cache := gcache.New()
	adap := adapter.NewRedis(g.Redis("zhxu"))
	cache.SetAdapter(adap)
	// change database cache from IN_memory to redis
	newadap := adapter.NewRedis(g.Redis("new"))
	g.DB().GetCache().SetAdapter(newadap)
	fmt.Println("===============================================")

}
