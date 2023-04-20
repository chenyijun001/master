package settings

import (
	"GinBlog/global"
	"GinBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 接受参数信息

type SettingUri struct {
	Name string `uri:"name"`
}

func (SettingApi) SettingInfoView(c *gin.Context) {
	var su SettingUri
	err := c.ShouldBindUri(&su)
	if err != nil {
		res.FailWithCode(res.ParameterError, c)
		return
	}
	switch su.Name {
	case "site_info":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	case "qiniu":
		res.OKWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMessage(fmt.Sprintf("无法修改关于%s的配置", su.Name), c)
		return
	}
}
