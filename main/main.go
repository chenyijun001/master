package main

import (
	"GinBlog/core"
	"GinBlog/flags"
	"GinBlog/global"
	"GinBlog/routers"
)

// @title
// @version	1.0
// @description	API文档描述
// @host	127.0.0.1:8090
// @BasePath /
func main() {
	core.InitCore()
	//命令行参数绑定  如果不输入go run main/main.go -db 是不会执行此函数方法的
	option := flags.Parse() //判断是否携带 -db 标识信息
	if flags.IsWebStop(option) {
		flags.SwitchOption(option) //如果携带 -db 就执行db中的函数方法
		return
	}
	r := routers.InitRouter()
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		global.Logs.Fatal("已经链接：", err.Error())
	}
}
