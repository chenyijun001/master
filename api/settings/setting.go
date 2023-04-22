package settings

import (
	"GinBlog/config"
	"GinBlog/global"
	"GinBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SettingApi struct {
}

// IsYaml 判断参数属于哪个文件
func IsYaml(c *gin.Context, su SettingUri) {
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
}
