package svc

import (
	"fmt"

	"cs5296-project/server/config"
	"cs5296-project/server/utils/log"
)

type ServiceContext struct {
}

func MustNewServiceContext(c *config.Config) *ServiceContext {
	var logLevel string
	if c.DebugMode {
		logLevel = "debug"
	} else {
		logLevel = "info"
	}
	err := log.InitLogger(c.LogPath, c.AppName, logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	return &ServiceContext{}
}
