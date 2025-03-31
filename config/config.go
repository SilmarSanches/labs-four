package config

import (
	"time"

	"github.com/spf13/viper"
)

type AppSettings struct {
	Port              string
	RedisAddr         string
	RedisPassword     string
	RedisDB           string
	RateLimitType     string
	DefaultTokenLimit int
	DefaultIPLimit    int
	RateLimitDuration time.Duration
	BlockDuration     time.Duration
}

func ProvideConfig() *AppSettings {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadConfig() (*AppSettings, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	appConfig := &AppSettings{
		Port:              viper.GetString("PORT"),
		RedisAddr:         viper.GetString("REDIS_ADDR"),
		RedisPassword:     viper.GetString("REDIS_PASSWORD"),
		RedisDB:           viper.GetString("REDIS_DB"),
		RateLimitType:     viper.GetString("RATE_LIMIT_TYPE"),
		DefaultTokenLimit: viper.GetInt("DEFAULT_TOKEN_LIMIT"),
		DefaultIPLimit:    viper.GetInt("DEFAULT_IP_LIMIT"),
		RateLimitDuration: time.Duration(viper.GetInt("RATE_LIMIT_DURATION")) * time.Second,
		BlockDuration:     time.Duration(viper.GetInt("BLOCK_DURATION")) * time.Second,
	}

	return appConfig, nil
}
