package main

import (
	"fmt"
	_ "go-gf-demo/router"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/gtcp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
)

func main() {
	go gtcp.NewServer("127.0.0.1:7100", func(c *gtcp.Conn) {
		defer c.Close()
		for {
			data, err := c.Recv(-1)
			if len(data) > 0 {
				// if err := c.Send(append([]byte(">\n"), data...)); err != nil {
				// 	fmt.Println(err)
				// }
				fmt.Println(string(data))
			}
			if err != nil {
				break
			}
		}
	}).Run()
	<-time.After(time.Second * 2)

	conn, err := gtcp.NewConn("127.0.0.1:7100")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		if err := conn.Send([]byte(gconv.String(i))); err != nil {
			glog.Error(err)
		}
		<-time.After(time.Second)
	}
	g.Wait()
}
