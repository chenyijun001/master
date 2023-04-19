package ctype

import (
	"database/sql/driver"
	"strings"
)

//创建公共数组

type Array []string

//创建scan输出方法

func (a *Array) Scan(value interface{}) error {
	v := value.([]byte)
	if string(v) == "" {
		*a = []string{}
		return nil
	}
	*a = strings.Split(string(v), "\n")
	return nil
}

//数字转换为值

func (a Array) Value() (driver.Value, error) {
	return strings.Join(a, "\n"), nil
}
