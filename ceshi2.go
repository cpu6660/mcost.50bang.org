package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

//mysql 配置格式
type MysqlConfig struct {
	Host string
	Port int
	User string
	Password string
}

//redis 配置格式
type RedisConfig struct {
	Host string
	Port string
	Auth string
}

//db 文件的配置格式
type DbConfig struct {
     Mysql MysqlConfig
     Redis RedisConfig
     Storage string
}

type JsonStruct struct {

}


func(self *JsonStruct) Load(filename string,v interface{}){

	//读取文件中的配置
	data,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data,v)
	if err != nil {
		fmt.Println(err)
	}
}



// 这些数据可以用类型模板表示  解释成模板  通过模板相应的操作获取到想要的数据
func main() {

	//var c DbConfig
	//var jsonParser JsonStruct
	//jsonParser.Load("./config/db.json",&c)
	//fmt.Println(c.Mysql)

	type Hconfig struct {
		Store string
		StoreConfig *json.RawMessage
	}

	var a  string = `{"Store" : "redis", "StoreConfig" : {"port" : "3306","host" : "127.0.0.1" }}`

	c,err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	var d Hconfig
	json.Unmarshal(c,&d)

	fmt.Println(d.Store)
}