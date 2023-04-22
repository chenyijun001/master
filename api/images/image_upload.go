package images

import (
	"GinBlog/global"
	"GinBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

//API上传

func (ImageApi) ImageUploadView(c *gin.Context) {
	formData, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	//获取上传图片
	fileList, ok1 := formData.File["images"]
	if !ok1 {
		res.FailWithMessage("file type not find", c)
		return
	}
	//判断文件是否存在，不存在则创建
	err = FilePathIfNotExist(global.Config.FileSettings.Path)
	if err != nil {
		res.FailWithCode(res.SettingError, c)
		return
	}
	for _, file := range fileList {
		path := global.Config.FileSettings.Path + "/" + file.Filename //获取文件路径及文件名称
		//判断文件大小及文件类型正确性
		ok2 := IsFileSize(file)
		if !ok2 {
			res.FailWithMessage(fmt.Sprintf("上传文件大小错误，文件大小不能大于%vMB", global.Config.FileSettings.Size), c)
			return
		}
		ok3 := IsFileSuffix(file)
		//ok3 := ValidateImageExt(file)
		if !ok3 {
			res.FailWithMessage("上传文件错误，文件类型不符合要求！", c)
			return
		}
		//开始真正上传文件
		err := SaveFile(c, file, path)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
	}
}
