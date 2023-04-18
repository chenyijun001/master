package main

import (
	"GinBlog/core"
	"GinBlog/global"
	"GinBlog/routers"
)

func main() {
	core.InitCore()
	r := routers.InitRouter()
	r.Run(global.Config.System.Addr())
}
