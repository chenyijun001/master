package models

import (
	"GinBlog/global"
	"GinBlog/models/ctype"
	"gorm.io/gorm"
	"os"
)

// 创建图片表

type BannerModel struct {
	MODEL
	Path    string          `json:"path"`                  //图片路径
	HashKey string          `json:"hash_key"`              //图片的hash值，用来判断图片是否为重复的
	Name    string          `gorm:"size:38" json:"name"`   //图片的名称
	Type    ctype.ImageType `gorm:"default:1" json:"type"` //开始事务的类型
}

// 拓展方法，增加事务

func (b BannerModel) BeforeDelete(db *gorm.DB) error {
	if b.Type == ctype.Local {
		//删除本地的缓存
		err := os.Remove(b.Path)
		if err != nil {
			global.Logs.Error(err)
			return err
		}
	}
	return nil
}
