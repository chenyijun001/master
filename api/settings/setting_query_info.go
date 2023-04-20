package settings

import (
	"GinBlog/global"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
