package models

// 创建文章标签

type TagModel struct {
	MODEL
	Title    string `gorm:"size:16" json:"title"`           //标签名称
	Articles string `gorm:"many2many:article_tag" json:"-"` //关联该标签的文章列表
}
