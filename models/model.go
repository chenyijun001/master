package models

// 编写通用字段表

type MODEL struct {
	ID       uint `gorm:"primaryKey" json:"id"` //主键ID
	CreateAt uint `json:"create_at"`            //创建时间
	UpdateAt uint `json:"-"`                    //更新时间
}
