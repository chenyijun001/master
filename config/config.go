package config

// 创建config结构体

type Config struct {
	Mysql        Mysql        `yaml:"mysql"`
	Logger       Logger       `yaml:"logger"`
	System       System       `yaml:"system"`
	SiteInfo     SiteInfo     `yaml:"site_info"`
	Email        Email        `yaml:"email"`
	QQ           QQ           `yaml:"qq"`
	Jwt          Jwt          `yaml:"jwt"`
	QiNiu        QiNiu        `yaml:"qiniu"`
	FileSettings FileSettings `yaml:"file_settings"`
}
