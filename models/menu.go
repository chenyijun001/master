package models

import "GinBlog/models/ctype"

//创建菜单表

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:36" json:"menu_title"`                                                              //中文标题
	MenuTitleEn  string        `gorm:"size:36" json:"menu_title_en"`                                                           //英文标题
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                  //标语
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                            //简介
	AbstractTime int           `json:"abstract_time"`                                                                          //简介的切换时间
	Banner       []BannerModel `gorm:"many2many:menu_image_models;joinForeignKey:MenuID;joinReferences:ImageID" json:"banner"` //菜单的图片列表
	BannerTime   int           `json:"banner_time"`                                                                            //菜单的切换时间
	Sort         int           `gorm:"size:10" json:"sort"`                                                                    //菜单的顺序
}
