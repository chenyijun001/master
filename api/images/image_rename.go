package images

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageRename struct {
	ID      int    `json:"id" binding:"required" msg:"请选择文件id"`       //根据ID修改名称
	NewName string `json:"new_name" binding:"required" msg:"请输入文件名称"` //输入所修改的新名称
}

func (ImageApi) ImageRenameView(c *gin.Context) {
	var imageNewName ImageRename
	err := c.ShouldBindJSON(&imageNewName)
	//校验绑定参数是否错误
	if err != nil {
		res.FailWithObjMsg(err, &imageNewName, c)
		global.Logs.Error("修改名称失败", err)
		return
	}
	//判断文件是否存在
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, imageNewName.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	//开始修改文件
	err = global.DB.Model(&imageModel).Update("name", imageNewName.NewName).Error
	if err != nil {
		res.FailWithMessage("文件名称修改失败"+err.Error(), c)
		return
	}
	res.OKWithMessage("文件修改成功", c)
	return
}
