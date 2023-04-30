package ctype

import "encoding/json"

//创建角色权限类型

type ImageType int

//创建角色常量

const (
	Local ImageType = 1 //本地数据库
	QiNiu ImageType = 2 //七牛云数据库存储空间
)

//创建Json解析

func (types ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(types.String())
}

func (types ImageType) String() string {
	var imageType string
	switch types {
	case 1:
		imageType = "本地"
	case 2:
		imageType = "七牛云"
	default:
		imageType = "其他"
	}
	return imageType
}
