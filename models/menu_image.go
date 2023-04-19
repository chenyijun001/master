package models

// 创建自定义菜单，方便排序 并且将自定义菜单和背景图的链接进行结合

type MenuImageModel struct {
	MenuID      string      `json:"menu_id"`             //菜单ID
	MenuModel   MenuModel   `gorm:"foreignkey:MenuID"`   //菜单列表
	BannerID    uint        `json:"banner_id"`           //图片id
	BannerModel BannerModel `gorm:"foreignkey:BannerID"` //图片列表
	Sort        int         `gorm:"size:10" json:"sort"` //菜单的顺序
}
