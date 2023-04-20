package conf

import (
	"embed"
	"time"
)

//go:embed *.yaml
var ConfYamlDir embed.FS

type AppConf struct {
	AppName string    `yaml:"app-name"`
	Server  ServerCfg `yaml:"server"`
	Redis   RedisCfg  `yaml:"redis"`
	Mysql   MysqlCfg  `yaml:"mysql"`
	Log     LogCfg    `yaml:"log"`
}

type ServerCfg struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     int           `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type MysqlCfg struct {
	Base struct {
		MaxOpenConn     int           `yaml:"maxOpenConn"`
		MaxIdleConn     int           `yaml:"maxIdleConn"`
		ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
	} `yaml:"base"`

	Read struct {
		Addr string `yaml:"addr"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"read"`

	Write struct {
		Addr string `yaml:"addr"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"write"`
}

type RedisCfg struct {
	Addr        string `yaml:"addr"`
	Db          int    `yaml:"db"`
	MaxRetries  int    `yaml:"maxRetries"`
	MinIdleConn int    `yaml:"minIdleConn"`
	Pass        string `yaml:"pass"`
	PoolSize    int    `yaml:"poolSize"`
}

type LogCfg struct {
	RuntimeRootPath string `yaml:"runtimeRootPath"`
	LogSavePath     string `yaml:"logSavePath"`
	LogSaveName     string `yaml:"logSaveName"`
	LogFileExt      string `yaml:"logFileExt"`
	TimeFormat      string `yaml:"timeFormat"`
}
