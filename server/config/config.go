package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DebugMode bool
	AppName   string
	ListenOn  string
	MySQL     string
	LogPath   string

	ClientOfflineThreshold int
}

func MustLoad(configFile string, config *Config) {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configFile)

	if err = v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	if err = v.Unmarshal(config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}
	return
}
