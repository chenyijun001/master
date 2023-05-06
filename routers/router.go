package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// 初始化路由
type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("api/")
	routerGroupApp := RouterGroup{apiGroup}
	routerGroupApp.SettingRouter()
	routerGroupApp.ImageRouter()
	return r
}
