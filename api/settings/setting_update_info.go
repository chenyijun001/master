package settings

import (
	"GinBlog/config"
	"GinBlog/core"
	"GinBlog/global"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingUpdateInfoView(c *gin.Context) {
	var si config.SiteInfo
	err := c.ShouldBindJSON(&si)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	global.Config.SiteInfo = si
	err = core.SetYaml()
	if err != nil {
		global.Logs.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	res.OKWithMessage("success update siteinfo", c)
}
