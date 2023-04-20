package core

import (
	"GinBlog/config"
	"GinBlog/global"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

//创建常量文件

const ConfigFile = "settings/settings.yaml" //取得config的地址

//InitConf 读取配置文件

func InitConf() {
	log.Println("<<<<<<<<<<<--------------------------config yaml start------------------------>>>>>>>>")
	conf := &config.Config{}
	global.Config = conf
	yamlConf, err := ioutil.ReadFile(ConfigFile) //读取配置文件中的信息
	if err != nil {
		panic(fmt.Errorf("get yamlConfig error:%s", err))
	}
	err = yaml.Unmarshal(yamlConf, conf) //格式化
	if err != nil {
		global.Logs.Fatalf("config init unmarshal error:%v", err)
	}
	log.Println("<<<<<<<<<<<--------------------------config yaml successful------------------------>>>>>>>>")
}
