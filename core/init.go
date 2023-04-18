package core

import "GinBlog/global"

//将所有初始化准备，在此处进行。

func InitCore() {
	InitConf()
	global.Logs = InitLogger()
	global.DB = InitMysql()
}
