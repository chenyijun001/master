package routers

import (
	"github.com/gin-gonic/gin"
)

// 初始化路由
type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	apiGroup := r.Group("api/")
	routerGroupApp := RouterGroup{apiGroup}
	routerGroupApp.SettingRouter()
	routerGroupApp.ImageRouter()
	return r
}
