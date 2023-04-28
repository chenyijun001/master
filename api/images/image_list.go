package images

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)



func (ImageApi) ImageListView(c *gin.Context) {
	//绑定参数
	var pg models.Page
	err := c.ShouldBindQuery(&pg)
	if err != nil {
		res.FailWithMessage("分页查询失败，请联系管理员", c)
		return
	}

	//创建
	var imageList []models.BannerModel

	//数据库查询总条数
	count := global.DB.Debug().Select("id").Find(&imageList).RowsAffected

	offset := (pg.Page - 1) * pg.Limit
	if offset < 0 {
		offset = 0
	}
	fmt.Println(pg)
	//进行分页
	global.DB.Debug().Limit(pg.Limit).Offset(offset).Find(&imageList)

	//设置编码返回
	res.OKWithList(imageList,count,c)

	return
}
