package config


var Configuration map[string]interface{}


//  无论是配置  还是读取  都很麻烦   确定类型
//  fmt.Println(config.Configuration["mysql"].(map[string]string)["host"] )
func init() {
	Configuration = make(map[string]interface{})

	//string 类型的
	Configuration["config1"] =  "value1"
	// int 类型的
	Configuration["config2"] =  12

	//map 类型的
	Configuration["mysql"] = make(map[string]string)
	Configuration["mysql"].(map[string]string)["host"] = "127.0.0.1"
	Configuration["mysql"].(map[string]string)["port"] = "3306"
	Configuration["mysql"].(map[string]string)["user"] = "root"
	Configuration["mysql"].(map[string]string)["password"] = "ceshi123"



}
