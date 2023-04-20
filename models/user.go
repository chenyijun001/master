package models

import "GinBlog/models/ctype"

// 用户表

type UserModel struct {
	MODEL
	NickName          string           `gorm:"size:36" json:"nick_name"`                                                              //用户昵称
	UserName          string           `gorm:"size:36" json:"user_name"`                                                              //用户名
	Password          string           `gorm:"size:128" json:"password"`                                                              //用户密码
	Avatar            string           `gorm:"size:36" json:"avatar"`                                                                 //用户头像
	Email             string           `gorm:"size:128" json:"email"`                                                                 //用户邮箱
	Tel               string           `gorm:"size:18" json:"tel"`                                                                    //用户电话
	Address           string           `gorm:"size:64" json:"address"`                                                                //用户地址
	Token             string           `gorm:"size:64" json:"token"`                                                                  //用户其他平台唯一ID
	IP                string           `gorm:"size:20" json:"ip"`                                                                     //用户ip地址
	Role              ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                          //用户权限：1 管理员 2 普通用户 3 游客 4 禁用用户
	SignStatus        ctype.SignStatus `gorm:"ctype=smallint(6)" json:"sign_status"`                                                  //用户注册来源
	ArticleModels     []ArticleModel   `gorm:"foreignkey:UserID" json:"-"`                                                            //用户发布文章列表
	UserCollectModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"` //用户收藏文章列表
}
