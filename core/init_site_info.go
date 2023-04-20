package core

import (
	"GinBlog/global"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
)

func SetYaml() error {
	data, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, data, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Logs.Info("success update yaml--siteInfo")
	return nil
}
