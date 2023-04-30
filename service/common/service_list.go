package common

import (
	"GinBlog/global"
	"GinBlog/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	//数据库查询总条数
	count = global.DB.Debug().Select("id").Find(&list).RowsAffected

	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	if option.Limit < 1 {
		option.Limit = 1
	}
	if option.Sort == "" {
		option.Sort = "asc"
	}
	//进行分页
	err = global.DB.Debug().Order("id " + option.Sort).Limit(option.Limit).Offset(offset).Find(&list).Error

	return list, count, err
}
