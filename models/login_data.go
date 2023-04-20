package models

//创建登陆信息表

type LoginDataModel struct {
	MODEL
	UserID    uint      `json:"user_id"`                             //用户ID
	UserModel UserModel `gorm:"foreignkey:UserID" json:"user_model"` //用户列表
	IP        string    `gorm:"size:20" json:"ip"`                   //用户IP
	NickName  string    `gorm:"size:42" json:"nick_name"`            //用户昵称
	Token     string    `gorm:"size:256" json:"token"`               //用户唯一身份认证
	Device    string    `gorm:"size:256" json:"device"`              //用户登录设备
	Address   string    `gorm:"size:64" json:"address"`              //用户登录地址
}
