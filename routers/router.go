package routers

import "github.com/gin-gonic/gin"

//初始化路由

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "demo")
	})
	return r
}
