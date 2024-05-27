package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Driver   string `mapstructure:"driver"`
	User     string
	Password string
	DB       string `mapstructure:"db"`
	Timeout  string `mapstructure:"timeout"`
}

func injectDBSecret(env *viper.Viper, app *AppConfig) {
	app.DB.User = env.GetString("DB_USER")
	app.DB.Password = env.GetString("DB_PASSWORD")
}
