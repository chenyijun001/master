package settings

import (
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
	//判断文件名称
	IsYaml(c, su)
	//修改yaml文件
	err = core.SetYaml()
	if err != nil {
		global.Logs.Error(err)
		res.FailWithCode(res.ParameterError, c)
		return
	}
	res.OKWithMessage(fmt.Sprintf("修改yaml--%s成功", su.Name), c)
}
