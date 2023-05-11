package api

import (
	"GinBlog/api/images"
	"GinBlog/api/menus"
	"GinBlog/api/settings"
)

type ApiGroup struct {
	SettingApi settings.SettingApi
	ImageApi   images.ImageApi
	MenuApi    menus.MenuApi
}

//创建实例

var App = new(ApiGroup)
