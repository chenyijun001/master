package routers

import (
	"GinBlog/api"
)

var setting = api.App.SettingApi

func (r RouterGroup) SettingRouter() {
	r.GET("settingsInfo/:name", setting.SettingInfoView)
	r.POST("settingsUpdate/:name", setting.SettingUpdateInfoView)
}
