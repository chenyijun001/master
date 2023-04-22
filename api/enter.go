package api

import (
	"GinBlog/api/images"
	"GinBlog/api/settings"
)

type ApiGroup struct {
	SettingApi settings.SettingApi
	ImageApi   images.ImageApi
}

//创建实例

var App = new(ApiGroup)
