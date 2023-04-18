package config

//创建系统信息

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s System) Addr() string {
	return s.Host + ":" + s.Port
}
