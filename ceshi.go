package main

import (
	"fmt"
	"gopkg.in/orivil/service.v0"
)

type Container struct{
	service map[string]int
}



const (
	// service name
	//
	// 为了保证唯一性, 建议用 "packageName.structName"
	SvcFather = "service_test.Father"
	SvcChild = "service_test.Child"


	// service priority
	FatherName = "good father"
)


func (c *Container) Add(service string,id int){
	c.service[service] = id
}

func (c *Container) Get(service string) int {
	return c.service[service]
}

func New() *App {
	return new(App)
}

//struct 匿名字段 这里只能说明拥有了相同的属性和操作
type App struct {
	Container
}

var a func(c *Container)interface{}

// service Father
type Father struct {
	name string
}

// service Child dependent on service Father
type Child struct {
	father *Father
}

// father provider
var fatherProvider = func(c *service.Container) interface{} {
	return &Father{name: FatherName}
}

// child provider
var childProvider = func(c *service.Container) interface{} {
	father := c.Get(SvcFather).(*Father) // get dependency service
	return &Child{father: father} // injection
}

func main() {

	//c := new(Container)
  	//a = hello
	//d := a(c)
	////只要是interface 类型的 都需要断言成为具体的类型
	//fmt.Println(d.(string)+"ddd")

	//c := service.NewPublicContainer()
	//
	//// 2. add providers
	//c.Add(SvcFather, fatherProvider)
	//
	//c.Add(SvcChild, childProvider)
	//
	//fmt.Println(c)
	//// 3. get service
	//// 测试依赖注入(dependency injection)
	//child := c.Get(SvcChild).(*Child)
	//father := child.father  // father was injected
	//fmt.Println(father.name == FatherName)



}


func hello(c *Container) interface{} {
	return "this is a test"
}