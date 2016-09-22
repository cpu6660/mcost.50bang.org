//应用实例

package bootstrap

import (
	"sync"
)
type application struct {
	//应用的版本信息
	version string
	//配置目录
	configPath string
	//application 应该具有服务容器的功能
	Container

}

func (app *application) GetAppVersion() string{
	return app.version
}

//设置应用的配置文件夹
func(app *application) SetConfigPath(path string) {
	app.configPath = path
}

//获取应用的配置文件夹
func(app *application) GetConfigPath() string {
	return app.configPath
}

//初始化container服务
func(app *application) initContainer(){
	app.providers = make(map[string]func()interface{})
	app.rWLocker  = new(sync.RWMutex)
}

