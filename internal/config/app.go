package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
)

type AppConfig struct {
	DB DBConfig `mapstructure:"database"`
}

const (
	configPath = "./internal/config/yml"
	fileType   = "yml"
)

func NewAppConfig() (*AppConfig, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigType(fileType)
	err := loadAllConfigFiles(v)

	var app AppConfig
	err = unmarshalConfig(v, &app)
	err = injectSecretConfig(&app)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func loadAllConfigFiles(v *viper.Viper) error {
	files, err := filepath.Glob(configPath + "/*")
	if err != nil {
		return fmt.Errorf("filepath is wrong %w", filepath.ErrBadPattern)
	}
	for _, file := range files {
		base := filepath.Base(file)
		v.SetConfigName(base)
		err := v.MergeInConfig()
		if err != nil {
			return fmt.Errorf("failed to merge viper config file %w", err)
		}
	}
	return nil
}

func injectSecretConfig(app *AppConfig) error {
	env := viper.New()
	env.AddConfigPath(".")
	env.SetConfigName(".env")
	env.SetConfigType("env")
	err := env.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read env file %w", err)
	}
	injectDBSecret(env, app)
	return nil
}

func unmarshalConfig(v *viper.Viper, target any) error {
	err := v.Unmarshal(&target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config %w", err)
	}
	return nil
}
