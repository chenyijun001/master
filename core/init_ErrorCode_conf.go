package core

import (
	"GinBlog/global"
	"GinBlog/models/res"
	"encoding/json"
	"os"
)

const file = "models/res/code.json"

type ErrMap map[res.ErrorCode]string

func InitError() *ErrMap {
	file, err := os.ReadFile(file)
	if err != nil {
		global.Logs.Fatalf("open code.json file error:%s", err)
	}
	ErrMsg := &ErrMap{}
	err = json.Unmarshal(file, ErrMsg)
	if err != nil {
		global.Logs.Fatalf("json resolve failed error:%s", err)
	}
	return ErrMsg
}
