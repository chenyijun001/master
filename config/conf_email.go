package config

type Email struct {
	Host             string `yaml:"level"`                                        //IP地址
	Port             int    `yaml:"port"`                                         //端口号
	User             string `yaml:"user"`                                         //用户
	Password         string `yaml:"password" json:"password"`                     //密码
	DefaultFromEmail string `yaml:"default_from_email" json:"default_from_email"` //发件人名称
	UseSSL           bool   `yaml:"use_ssl" json:"use_ssl"`                       //SSL链接
	UserTLS          bool   `yaml:"user_tls" json:"user_tls"`                     //TLS链接
}
