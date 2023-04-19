package models

import "time"

//创建收藏表：记录用户何时收藏何文章

type CollectsModel struct {
	UserID       uint         `gorm:"primaryKey"`           //用户ID
	UserModel    string       `gorm:"foreignkey:UserID"`    //关联的用户
	ArticleID    uint         `gorm:"primaryKey"`           //文章ID
	ArticleModel ArticleModel `gorm:"foreignkey:ArticleID"` //关联的文章
	CreateAt     time.Time    //创建时间
}
