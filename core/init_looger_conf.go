package core

import (
	"GinBlog/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

//配置logger日志的初始化

//配置颜色

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
}

//format实现 formatter 方法

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	//根据不同等级，log展示不同颜色
	var levelColor int

	switch entry.Level { //根据不同log日志等级，分配不同的显示颜色
	case logrus.DebugLevel, logrus.TraceLevel: //debug和trace等级显示为灰色
		levelColor = gray
	case logrus.WarnLevel: //warn等级显示为黄色
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel: //error、fatal、panic等级显示为红色
		levelColor = red
	default: //其他等级显示为蓝色
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//自定义日期格式
	timeStamp := entry.Time.Format("2023-01-01 12:01:01")

	//定义一个显示文件名log
	logs := global.Config.Logger
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s]\x1b[%dm[%s]\x1b[0m %s %s %s\n", logs.Prefix, timeStamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s]\x1b[%dm[%s]\x1b[0m %s \n", logs.Prefix, timeStamp, levelColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

//初始化日志

func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                //创建实例
	mLog.SetOutput(os.Stdout)                           //设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名称和行号
	mLog.SetFormatter(&LogFormatter{})                  //设置自己定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置最低log等级
	//修改全局log的日志
	InitDefaultLogger()
	return mLog
}

//初始化全局log

func InitDefaultLogger() {
	//全局log
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名称和行号
	logrus.SetFormatter(&LogFormatter{})                  //设置自己定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低log等级
}
