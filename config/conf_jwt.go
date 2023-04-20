package config

type Jwt struct {
	Secret  string `yaml:"secret" json:"secret"`   //秘钥
	Expires int    `yaml:"expires" json:"expires"` //过期时间
	Issuer  string `yaml:"issuer" json:"issuer"`   //颁发人
}
