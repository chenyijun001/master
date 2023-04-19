package models

// 创建图片表

type BannerModel struct {
	MODEL
	Path    string `json:"path"`                //图片路径
	HashKey string `json:"hash_key"`            //图片的hash值，用来判断图片是否为重复的
	Name    string `gorm:"size:38" json:"name"` //图片的名称
}
