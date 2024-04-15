package svc

import (
	"fmt"
	"os"

	"github.com/SunSetPilot/cs5296-project/server/config"
	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/utils/log"
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

	if os.Getenv("MYSQL_DSN") != "" {
		c.MySQL = os.Getenv("MYSQL_DSN")
	}
	err = dal.InitDB(c.MySQL)
	if err != nil {
		panic(fmt.Errorf("failed to init db: %w", err))
	}
	return &ServiceContext{
		SvcConf: c,
	}
}
