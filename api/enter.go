package api

import "GinBlog/api/settings"

type ApiGroup struct {
	SettingApi settings.SettingApi
}

//创建实例
var App = new(ApiGroup)
