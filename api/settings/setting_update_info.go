package settings

import (
	"GinBlog/config"
	"GinBlog/core"
	"GinBlog/global"
	"GinBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingUpdateInfoView(c *gin.Context) {
	var su SettingUri
	err := c.ShouldBindUri(&su)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	switch su.Name {
	case "site_info":
		var si config.SiteInfo
		err := c.ShouldBindJSON(&si)
		if err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.SiteInfo = si
	case "qq":
		var qq config.QQ
		err := c.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.QQ = qq
	case "email":
		var email config.Email
		err := c.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.Email = email
	case "jwt":
		var jwt config.Jwt
		err := c.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.Jwt = jwt
	case "qiniu":
		var qiniu config.QiNiu
		err := c.ShouldBindJSON(&qiniu)
		if err != nil {
			res.FailWithCode(res.ParameterError, c)
			return
		}
		global.Config.QiNiu = qiniu
	default:
		res.FailWithMessage(fmt.Sprintf("无法修改关于%s的配置", su.Name), c)
		return
	}
	err = core.SetYaml()
	if err != nil {
		global.Logs.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	res.OKWithMessage(fmt.Sprintf("修改yaml--%s成功", su.Name), c)
}
