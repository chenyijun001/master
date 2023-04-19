package models

//创建消息表，记录消息

type MessageModel struct {
	MODEL
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"`               //发送人ID
	SendUserModel    UserModel `gorm:"foreignkey:SendUserID" json:"send_user_model"` //发送人列表
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"`           //发送人昵称
	SendUserAvatar   string    `json:"send_user_avatar"`                             //发送人头像
	RevUserID        uint      `gorm:"primaryKey" json:"rev_user_id"`                //接收人ID
	RevUserModel     UserModel `gorm:"foreignkey:RevUserID" json:"rev_user_model"`   //接收人列表
	RevUserNickName  string    `gorm:"size:42" json:"rev_user_nick_name"`            //接收人昵称
	RevUserAvatar    string    `json:"rev_user_avatar"`                              //发送人头像
	IsSend           string    `gorm:"default:false" json:"is_send"`                 //接收方是否查看
	Content          string    `json:"content"`                                      //消息内容
}
