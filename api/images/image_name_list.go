package images

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

func (ImageApi) ImageNameListView(c *gin.Context) {
	var list []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&list)
	//设置编码返回
	res.OKWithData(list, c)
}
