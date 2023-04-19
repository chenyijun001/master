package models

import "GinBlog/models/ctype"

// 编写通用字段表

type MODEL struct {
	ID       uint `gorm:"primaryKey" json:"id"` //主键ID
	CreateAt uint `json:"create_at"`            //创建时间
	UpdateAt uint `json:"-"`                    //更新时间
}

// 用户表

type UserModel struct {
	NickName       string           `gorm:"size:36" json:"nick_name"`                         //用户昵称
	UserName       string           `gorm:"size:36" json:"user_name"`                         //用户名
	Password       string           `gorm:"size:128" json:"password"`                         //用户密码
	AvatarID       uint             `json:"avatar_id"`                                        //用户头像ID
	Avatar         string           `gorm:"size:36" json:"avatar"`                            //用户头像
	Email          string           `gorm:"size:128" json:"email"`                            //用户邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                               //用户电话
	Address        string           `gorm:"size:64" json:"address"`                           //用户地址
	Token          string           `gorm:"size:64" json:"token"`                             //用户其他平台唯一ID
	Ip             string           `gorm:"size:20" json:"ip"`                                //用户ip地址
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`                     //用户权限：1 管理员 2 普通用户 3 游客 4 禁用用户
	SignStatus     ctype.SignStatus `gorm:"ctype=smallint(6)" json:"sign_status"`             //用户注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignkey:UserID" json:"-"`                       //用户发布文章列表
	CollectsModels []CollectsModel  `gorm:"many2many:auth2_collects;joinForeignKey" json:"-"` //用户收藏文章列表
}
