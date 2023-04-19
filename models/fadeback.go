package models

//创建用户反馈表

type FadeBackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`          //邮箱
	Content      string `gorm:"size:128" json:"content"`       //内容
	ApplyContent string `gorm:"size:128" json:"apply_content"` //回复内容
	IsApply      bool   `json:"is_apply"`                      //是否回复
}
