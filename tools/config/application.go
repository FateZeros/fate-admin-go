package config

import "github.com/spf13/viper"

type Application struct {
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	Mode          string
	DemoMsg       string
	Domain        string
	IsHttps       bool
}

var ApplicationConfig = new(Application)

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		ReadTimeout:   cfg.GetInt("readTimeout"),
		WriterTimeout: cfg.GetInt("writerTimeout"),
		Host:          cfg.GetString("host"),
		Port:          portDefault(cfg),
		Name:          cfg.GetString("name"),
		JwtSecret:     cfg.GetString("jwtSecret"),
		Mode:          cfg.GetString("mode"),
		DemoMsg:       cfg.GetString("demoMsg"),
		Domain:        cfg.GetString("domain"),
		IsHttps:       cfg.GetBool("ishttps"),
	}
}

func portDefault(cfg *viper.Viper) string {
	if cfg.GetString("port") == "" {
		return "8000"
	} else {
		return cfg.GetString("port")
	}
}
