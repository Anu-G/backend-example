package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBConfig
	APIConfig
}

type DBConfig struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	DBPort      string `mapstructure:"DB_PORT"`
	SSLMode     string `mapstructure:"SSL_MODE"`
	TimeZone    string `mapstructure:"TIME_ZONE"`
	Environment string `mapstructure:"ENV"`
}

type APIConfig struct {
	APIUrl string `mapstructure:"API_URL"`
}

func (c *Config) loadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	if err = viper.Unmarshal(&config.APIConfig); err != nil {
		return
	}

	if err = viper.Unmarshal(&config.DBConfig); err != nil {
		return
	}
	return
}

func NewConfig() Config {
	cfg := Config{}
	cfg, err := cfg.loadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return cfg
}
