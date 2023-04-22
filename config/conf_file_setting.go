package config

import "GinBlog/models/ctype"

//创建文件配置

type FileSettings struct {
	Path string      `yaml:"path" json:"path"` //文件保存地址
	Size float64     `yaml:"size" json:"size"` //文件上传限制大小
	Type ctype.Array `yaml:"type" json:"type"` //文件图片限制类型
}
