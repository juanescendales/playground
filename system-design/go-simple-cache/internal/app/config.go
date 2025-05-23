package app

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

var ErrNoConfigDir = errors.New("CONFIG_DIR environment variable is not set")

type Config struct {
	CacheCapacity int `mapstructure:"cache-capacity"`
}

func LoadConfig() (Config, error) {
	var config Config
	configDirectory, exists := os.LookupEnv("CONFIG_DIR")
	if !exists {
		return config, ErrNoConfigDir
	}
	configExtractor := viper.New()

	configExtractor.SetConfigName("config")
	configExtractor.SetConfigType("yml")
	configExtractor.AddConfigPath(configDirectory)
	err := configExtractor.ReadInConfig()
	if err != nil {
		return config, err
	}

	err = configExtractor.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
