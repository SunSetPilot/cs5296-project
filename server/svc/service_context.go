package svc

import (
	"cs5296-project/server/table"
	"cs5296-project/utils/log"

	"fmt"

	"cs5296-project/server/config"
)

type ServiceContext struct {
	SvcConf *config.Config
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

	err = table.InitDB(c.MySQL)
	if err != nil {
		panic(fmt.Errorf("failed to init db: %w", err))
	}
	return &ServiceContext{
		SvcConf: c,
	}
}
