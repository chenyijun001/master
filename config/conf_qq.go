package config

import "fmt"

type QQ struct {
	AppID    string `yaml:"app_id" json:"app_id"`     //认证ID
	Key      string `yaml:"key" json:"key"`           //QQkey
	Redirect string `yaml:"redirect" json:"redirect"` //回调地址
}

// 设置跳转的地址
func (q QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_url=%s", q.AppID, q.Redirect)
}
