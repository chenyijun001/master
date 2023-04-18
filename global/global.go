package global

import (
	"GinBlog/config"
	"GinBlog/core"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//保存全局配置变量
//可以获取所有的配置信息

var (
	Config *config.Config
	DB     *gorm.DB
	Logs   *logrus.Logger
	ErrMsg *core.ErrMap
)
