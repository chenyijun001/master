package images

import (
	"GinBlog/global"
	"GinBlog/models"
	"GinBlog/models/res"
	"GinBlog/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type ImageApi struct {
}

// SaveFile 保存文件函数
func SaveFile(c *gin.Context, file *multipart.FileHeader, path string) error {
	hashKey, err1 := utils.MD5ByByte(file)
	if err1 != nil {
		return err1
	}
	//保存文件
	err := c.SaveUploadedFile(file, path)
	if err != nil {
		//此处后续可以补充一个判断图片类型日志输出，可以检测是否有黑客进行渗透
		return err
	}
	//判断该图片是否已经在数据库存在
	var banner models.BannerModel
	row := global.DB.Take(&banner, "hash_key=?", hashKey).RowsAffected
	if row != 0 {
		return errors.New("图片已经存在！")
	}
	//将图片入库
	global.DB.Create(&models.BannerModel{
		Path:    path,
		HashKey: hashKey,
		Name:    file.Filename,
	})
	global.Logs.Infof("上传文件成功,文件名为：%s,文件大小为：%v,文件header为：%s", file.Filename, file.Size, file.Header)
	res.OKWithMessage("上传成功！", c)
	return nil
}

// IsFileSize 判断文件大小函数
func IsFileSize(file *multipart.FileHeader) bool {
	size := math.Round((float64(file.Size)/1024/1024)*100) / 100
	//判断文件大小是否正确
	if size > global.Config.FileSettings.Size {
		global.Logs.Warnf("上传文件失败，上传数据大小超过限制,文件名为：%s,文件大小为：%v,文件header为：%s", file.Filename, file.Size, file.Header)
		return false
	}
	return true
}

// IsFileSuffix 判断文件类型是否正确
//func IsFileSuffix(file *multipart.FileHeader) bool {
//	filename := file.Filename
//	suffix := strings.SplitAfterN(filename, ".", 2)
//	s := suffix[1]
//	fmt.Println(global.Config.FileSettings.Type[1])
//	if s != "jpg" && s != "png" && s != "jpeg" && s != "gif" && s != "ico" && s != "tiff" && s != "webp" {
//		return false
//	}
//	return true
//}

// IsFileSuffix 判断文件类型是否正确 2.0 可自定义配置
func IsFileSuffix(file *multipart.FileHeader) bool {
	filename := file.Filename
	suffix := strings.SplitAfterN(filename, ".", 2)
	s := suffix[1]
	imageSuffix := global.Config.FileSettings.Type
	for _, imageType := range imageSuffix {
		if s == imageType {
			return true
		}
	}
	return false
}

// ValidateImageExt 校验图片后缀第二种方式 ---此方法存在bug，无法校验图片后缀的真实性
func ValidateImageExt(file *multipart.FileHeader) bool {
	filename := file.Filename
	ext := filepath.Ext(filename)
	fmt.Println("------------------------>>>>>>>", ext)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true // 合法的后缀名
	default:
		return false // 非法的后缀名
	}
}

// FilePathIfNotExist 函数检查指定路径是否存在，如果路径不存在，则使用 os.MkdirAll 创建路径及其所有必需的父文件夹。
// os.ModePerm 参数指定新创建的目录的权限
func FilePathIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
