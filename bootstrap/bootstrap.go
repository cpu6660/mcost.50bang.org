//项目启动文件
package bootstrap

var App *application
var logger func()interface{}
var config func()interface{}

//logger实现
func logfunc() interface{} {
	return "logger"
}

//config 实现
func confunc() interface{} {
	return "config"
}


//项目运行初始化
func init() {

	App = new(application)
	//初始化容器服务
        App.initContainer()

	//设置日志服务
	logger  = logfunc
	App.Add("logger",logger)

	//设置配置服务
	config = confunc
	App.Add("config",config)

}



