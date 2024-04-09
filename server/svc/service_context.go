package svc

import (
	"cs5296-project/server/model"
	"fmt"

	"cs5296-project/server/config"
	"cs5296-project/server/utils/log"
)

type ServiceContext struct {
}

func MustNewServiceContext(c *config.Config) *ServiceContext {
	var (
		logLevel string
		err      error
	)
	if c.DebugMode {
		logLevel = "debug"
	} else {
		logLevel = "info"
	}
	err = log.InitLogger(c.AppName, c.LogPath, logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	err = model.InitDB(c.MySQL)
	if err != nil {
		panic(fmt.Errorf("failed to init db: %w", err))
	}
	return &ServiceContext{}
}
