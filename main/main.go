package main

import (
	"GinBlog/core"
	"GinBlog/flags"
	"GinBlog/global"
	"GinBlog/routers"
	"fmt"
)

func main() {
	core.InitCore()

	//命令行参数绑定
	option := flags.Parse()
	fmt.Println(option)
	if flags.IsWebStop(option) {
		flags.SwitchOption(option)
		return
	}
	r := routers.InitRouter()
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		global.Logs.Fatal("已经链接：", err.Error())
	}
}
