package models

import (
	"gorm.io/gorm"
)

// 编写通用字段表

type MODEL struct {
	gorm.Model
}
type Page struct {
	Page  int    `form:"page"`  //页数
	Key   string `form:"key"`   //关键词
	Limit int    `form:"limit"` //每页限制多少条
	Sort  string `form:"sort"`  //排序
}