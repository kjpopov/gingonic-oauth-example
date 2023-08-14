package config

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type AppConfig struct {
	Oauth2Config oauth2.Config `mapstructure:"oauth2"`
}

// Global config variable accessible from all packages
var Config AppConfig

func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.BindEnv("oauth2.ClientID", "CLIENT_ID")
	viper.BindEnv("oauth2.Secret", "CLIENT_SECRET")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Config)
	return
}
