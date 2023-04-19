package ctype

import "encoding/json"

//创建角色权限类型

type SignStatus int

//创建角色常量

const (
	SignQQ     SignStatus = 1 //QQ
	SignWeChat SignStatus = 2 //微信
	SignGitee  SignStatus = 3 //Git
	SignEmail  SignStatus = 4 //Email
)

//创建Json解析

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var signType string
	switch s {
	case SignQQ:
		signType = "QQ"
	case SignWeChat:
		signType = "微信"
	case SignGitee:
		signType = "Git"
	case SignEmail:
		signType = "Email"
	default:
		signType = "其他"
	}
	return signType
}
