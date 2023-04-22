package models

import "time"

// 编写通用字段表

type MODEL struct {
	ID       uint      `gorm:"primaryKey" json:"id"` //主键ID
	CreateAt time.Time `json:"create_at"`            //创建时间
	UpdateAt time.Time `json:"-"`                    //更新时间
}
