package svc

import (
	"fmt"

	"github.com/SunSetPilot/cs5296-project/server/config"
	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

func NewMockServiceContext() *ServiceContext {
	var (
		logLevel string
		err      error
	)
	c := config.NewMockConfig()
	if c.DebugMode {
		logLevel = "debug"
	} else {
		logLevel = "info"
	}
	err = log.InitLogger(c.AppName, c.LogPath, logLevel)
	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	err = dal.InitMockDB()
	if err != nil {
		panic(fmt.Errorf("failed to init db: %w", err))
	}
	return &ServiceContext{
		SvcConf: c,
	}
}
