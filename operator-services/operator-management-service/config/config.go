package config

import (
	"github.com/spf13/viper"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Database DatabaseEnv `mapstructure:",squash"`
}

type DatabaseEnv struct {
	Uri string `mapstructure:"uri"`
}

func GetInstance() *Config {
	once.Do(func() {
		configLoader, err := loadConfig()
		if err != nil {
			panic(err)
		}
		cfg = configLoader
	})
	return cfg
}

func loadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
