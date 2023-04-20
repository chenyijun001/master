package flags

import (
	"GinBlog/global"
	"GinBlog/models"
)

//执行具体数据库命令

func MakeMigrations() {
	var err error
	global.DB.SetupJoinTable(&models.UserModel{}, "user_collect_models", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.BannerModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		&models.ArticleModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
		&models.UserCollectModel{},
	)

	if err != nil {
		global.Logs.Error("[error] 生成数据库表结构失败")
		return
	}
	global.Logs.Info("[success] 生成数据库表结构成功！")
}
