package main


import (
	"fmt"
	//"./config"
	"github.com/go-ini/ini"
)

const (
	LOG_PATH = "/user/log/"
	CONFIG_PATH  = "/user/config/"
)

var f *ini.File

func main() {

 // 1 fmt.Println(config.Configuration["mysql"].(map[string]string)["host"] )

 // 2 load ini

	session,err :=  f.GetSection("mysql")
	if err != nil {
		fmt.Println("error")
	}
	key1,err := session.GetKey("host")

	//取出以后,再转化为特定类型的值
	a := key1.String()
	fmt.Printf("%T",a)

}

func init() {

	//   如果直接写成这样,系统会将f 定义为函数内部的变量   f,err := ini.Load("./config/app.ini")
	var err error
	f,err = ini.Load("./config/app.ini")
	if err != nil {
		fmt.Println("error")
	}

}
