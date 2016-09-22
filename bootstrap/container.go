// container 服务容器

package bootstrap

import (
	"fmt"
	"sync"
)


type Container struct {
	//服务提供者
	providers map[string]func()interface{}
	//读写锁
	rWLocker *sync.RWMutex

}

func NewContainer() *Container {
	return &Container{
		providers : make(map[string]func()interface{},10),
		rWLocker  : new(sync.RWMutex),
	}
}

//添加服务
func (c *Container) Add(name string, provider func()interface{}){
	if _,ok := c.providers[name];ok{
		panic(fmt.Errorf("service.provider: service %s already registered! ",name))
	}

	c.providers[name] = provider
}

//获取相应的服务 get 给的是interface 类型,这里要注意一下,也可以给定一个规定的类型
func (c *Container) Get(service string)(instance interface{}){
	var ok bool
	//如果是只读的话,就通过,如果此时有写操作,就等待
	c.rWLocker.Lock()
	instance,ok = c.providers[service]
	c.rWLocker.Unlock()
	//如果存在 就返回该服务
	if ok {
		return
	}

	return nil
}