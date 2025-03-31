package config

import "github.com/spf13/viper"

type AppSettings struct {
	Port                        string
	RedisAddr                   string
	RedisPassword               string
	RedisDB                     string
	DefaultLimitPerSecond       string
	DefaultBlockDurationSeconds string
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

	viper.SetDefault("PORT", "")
	viper.SetDefault("API_KEY_TEMPO", "")

	appConfig := &AppSettings{
		Port:                        viper.GetString("PORT"),
		RedisAddr:                   viper.GetString("REDIS_ADDR"),
		RedisPassword:               viper.GetString("REDIS_PASSWORD"),
		RedisDB:                     viper.GetString("REDIS_DB"),
		DefaultLimitPerSecond:       viper.GetString("DEFAULT_LIMIT_PER_SECOND"),
		DefaultBlockDurationSeconds: viper.GetString("DEFAULT_BLOCK_DURATION_SECONDS"),
	}

	return appConfig, nil
}
