package images

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/res"
	"GinBlog/service/common"
	"github.com/gin-gonic/gin"
)

func (ImageApi) ImageListView(c *gin.Context) {
	//绑定参数
	var pg models.PageInfo
	err := c.ShouldBindQuery(&pg)
	if err != nil {
		res.FailWithMessage("分页查询失败，请联系管理员", c)
		return
	}

	////创建
	//var imageList []models.BannerModel
	//
	////数据库查询总条数
	//count := global.DB.Debug().Select("id").Find(&imageList).RowsAffected
	//
	//offset := (pg.Page - 1) * pg.Limit
	//if offset < 0 {
	//	offset = 0
	//}
	////进行分页
	//global.DB.Debug().Limit(pg.Limit).Offset(offset).Find(&imageList)

	//调用通用方法，进行传递list集合值
	list, count, err := common.CommonList(models.BannerModel{}, common.Option{
		pg,
		true,
	})
	if err != nil {
		global.Logs.Error(err)
	}
	//设置编码返回
	res.OKWithList(list, count, c)

	return
}
