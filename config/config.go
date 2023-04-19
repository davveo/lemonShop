package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type AppCfg struct {
	AppName   string
	PageSize  int
	PrefixUrl string
}

type ServerCfg struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DbCfg struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type RedisCfg struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type LogCfg struct {
	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

type Config struct {
	App    AppCfg
	Redis  RedisCfg
	DB     DbCfg
	Server ServerCfg
	Log    LogCfg
}

var (
	cfg  *ini.File
	Conf *Config
)

func Init() error {
	var err error

	Conf = &Config{
		App:    AppCfg{},
		Server: ServerCfg{},
		DB:     DbCfg{},
		Redis:  RedisCfg{},
		Log:    LogCfg{},
	}

	cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
		return err
	}

	mapTo("app", &Conf.App)
	mapTo("server", &Conf.Server)
	mapTo("database", &Conf.DB)
	mapTo("redis", &Conf.Redis)
	mapTo("log", &Conf.Log)

	Conf.Server.ReadTimeout = Conf.Server.ReadTimeout * time.Second
	Conf.Server.WriteTimeout = Conf.Server.WriteTimeout * time.Second
	Conf.Redis.IdleTimeout = Conf.Redis.IdleTimeout * time.Second

	return nil
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
