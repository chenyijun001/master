package settings

import (
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingApi) SettingInfoView(c *gin.Context) {
	//res.FailWithCode(res.SettingError, c)
	res.FailWithCode(res.SettingError, c)
}
