package flags

import sys_flag "flag"

//创建命令行脚手架

type Option struct {
	DB bool
}

// Parse 解析命令行参数

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	//解析命令行参数写入注册的flag里面
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

// 执行完数据库是否启动web

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

//根据不同的命令执行函数

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
}
