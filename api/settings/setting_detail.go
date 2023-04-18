package settings

import (
	_ "GinBlog/api"
	_ "GinBlog/config"
	_ "GinBlog/core"
	"GinBlog/global"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	res.FailWithCode(res.SettingError, c)
	res.FailWithMessage(
		global.ErrMsg[res.SettingError],
		c,
	)
}
