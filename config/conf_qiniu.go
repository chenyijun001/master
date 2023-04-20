package config

type QiNiu struct {
	AccessKey string  `yaml:"access_key" json:"access_key"` //秘钥key
	SecretKey string  `yaml:"secret_key" json:"secret_key"` //秘钥key
	Bucket    string  `yaml:"bucket" json:"bucket"`         //存储桶的名称
	CDN       string  `yaml:"cdn" json:"cdn"`               //访问图片的地址的编辑
	Zone      string  `yaml:"zone" json:"zone"`             //存储的地区
	Size      float64 `yaml:"size" json:"size"`             //允许存储的大小 单位是MB
}
