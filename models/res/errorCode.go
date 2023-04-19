package res

import (
	"GinBlog/global"
	"encoding/json"
	"os"
)

/*
	所有的错误码，都会放在这
*/

type ErrorCode int

const (
	SettingError ErrorCode = 1001 //系统错误
)

const file = "models/res/code.json"

type ErrMap map[ErrorCode]string

func InitError() ErrMap {
	file, err := os.ReadFile(file)
	if err != nil {
		global.Logs.Fatalf("open code.json file error:%s", err)
	}
	ErrMsg := &ErrMap{}
	err = json.Unmarshal(file, ErrMsg)
	if err != nil {
		global.Logs.Fatalf("json resolve failed error:%s", err)
	}
	return *ErrMsg
}
