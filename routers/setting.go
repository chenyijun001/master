package routers

import (
	"GinBlog/api"
)

var setting = api.App.SettingApi

func (r RouterGroup) SettingRouter() {
	r.GET("/settings", setting.SettingInfoView)
}
