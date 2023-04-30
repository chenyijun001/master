package images

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (ImageApi) DeleteImageByIdView(c *gin.Context) {
	var ids models.IDs
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		global.Logs.Error(err)
	}
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, ids.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("未删除任何数据", c)
		return
	}
	//global.DB.Unscoped().Delete(&imageList)	//此方法为永久删除
	global.DB.Unscoped().Delete(&imageList) //此方法软删除，只会删除本地服务器中的图片，不会删除服务器中记录
	res.OKWithData(&imageList, c)
}
