package core

import (
	"GinBlog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitMysql() *gorm.DB {
	DB := MysqlConnect()
	return DB
}
func MysqlConnect() *gorm.DB {
	log.Println("<<<<<<<<<<<--------------------------config mysql start------------------------>>>>>>>>")
	if global.Config.Mysql.Host == "" {
		//打印日志信息
		global.Logs.Warnln("connect mysql failed,host is null")
		return nil
	}
	//获取数据连接配置信息
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		logger.Default.LogMode(logger.Info)
	} else {
		logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Logs.Fatalf("[%s]connect mysql failed,error:%v", dsn, err)
	}
	sql, err := db.DB()
	sql.SetConnMaxLifetime(time.Hour * 4)
	sql.SetMaxOpenConns(100)
	sql.SetMaxIdleConns(10)
	log.Println("<<<<<<<<<<<--------------------------config mysql successful------------------------>>>>>>>>")
	return db
}
