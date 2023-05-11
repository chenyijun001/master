package routers

import "GinBlog/api"

var menu = api.App.MenuApi

func (r RouterGroup) MenuRouter() {
	r.POST("MenuCreate", menu.MenuCreateView)
}
