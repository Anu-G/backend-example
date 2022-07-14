package config

import (
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig
	APIConfig
	LopeiGrpcConfig
	TokenConfig
	RedisClient
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

type LopeiGrpcConfig struct {
	LopeiUrl string `mapstructure:"LOPEI_URL"`
}

type TokenConfig struct {
	ApplicationName     string `mapstructure:"APP_NAME"`
	JwtSignatureKey     string `mapstructure:"SECRET_KEY"`
	JwtSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
	Redis               *redis.Client
}

type RedisClient struct {
	RedisAddress string `mapstructure:"REDIS_ADDRESS"`
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

	if err = viper.Unmarshal(&config.LopeiGrpcConfig); err != nil {
		return
	}

	if err = viper.Unmarshal(&config.TokenConfig); err != nil {
		return
	}
	config.TokenConfig.JwtSigningMethod = jwt.SigningMethodHS256
	config.TokenConfig.AccessTokenLifeTime = 10 * time.Minute

	if err = viper.Unmarshal(&config.RedisClient); err != nil {
		return
	}
	newRedisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		DB:   0,
	})
	config.TokenConfig.Redis = newRedisClient
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
