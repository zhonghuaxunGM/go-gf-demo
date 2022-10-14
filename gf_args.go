package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gogf/gf/container/gvar"
)

var cmdOptions map[string]string

func init() {
	cmdOptions = make(map[string]string, 0)
}
func argsDemo() {
	reg := regexp.MustCompile(`\-\-{0,1}(.+?)=(.+)`)
	for i := 0; i < len(os.Args); i++ {
		result := reg.FindStringSubmatch(os.Args[i])
		fmt.Println(result)
		if len(result) > 1 {
			cmdOptions[result[1]] = result[2]
		}
	}
	fmt.Println(cmdOptions)
	fmt.Println(Get("otj").String())
	fmt.Println(Get("GOPATH").String())
	fmt.Println(Get("SSS").String(), "test")
	fmt.Println("===============================================")
}

func Get(key string, def ...interface{}) *gvar.Var {
	value := interface{}(nil)
	if len(def) > 0 {
		value = def[0]
	}
	cmdKey := strings.ToLower(strings.Replace(key, "_", ".", -1))
	if v, ok := cmdOptions[cmdKey]; ok {
		value = v
	} else {
		envKey := strings.ToUpper(strings.Replace(key, ".", "_", -1))
		if v := os.Getenv(envKey); v != "" {
			value = v
		}
	}
	return gvar.New(value)
}
