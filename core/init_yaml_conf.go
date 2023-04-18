package core

import (
	"GinBlog/config"
	"GinBlog/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

//InitConf 读取配置文件
func InitConf() {
	log.Println("<<<<<<<<<<<--------------------------config yaml start------------------------>>>>>>>>")
	//创建常量文件
	const ConfigFile = "settings/settings.yaml"
	//取得config的地址
	conf := &config.Config{}
	global.Config = conf
	//读取配置文件中的信息
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConfig error:%s", err))
	}
	//格式化
	err = yaml.Unmarshal(yamlConf, conf)
	if err != nil {
		global.Logs.Fatalf("config init unmarshal error:%v", err)
	}
	log.Println("<<<<<<<<<<<--------------------------config yaml successful------------------------>>>>>>>>")
}
