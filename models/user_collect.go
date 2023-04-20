package models

import "time"

//创建收藏表：记录用户何时收藏何文章

type UserCollectModel struct {
	UserID       uint         `gorm:"primaryKey"`           //用户ID
	UserModel    UserModel    `gorm:"foreignkey:UserID"`    //关联的用户
	ArticleID    uint         `gorm:"primaryKey"`           //文章ID
	ArticleModel ArticleModel `gorm:"foreignkey:ArticleID"` //关联的文章
	CreatedAt    time.Time    `json:"created_at"`           //创建时间
}
