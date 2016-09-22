package main

import (
	"fmt"
	"./bootstrap"
)

func main() {
	//设置config配置
	bootstrap.App.SetConfigPath("/hello/world/")
	c := bootstrap.App.Get("config")
	//因为设置的时候返回类型为interface,所以要断言
	fmt.Println(c.(func()interface{})())
}
