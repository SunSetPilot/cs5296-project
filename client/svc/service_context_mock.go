package svc

import (
	"fmt"
	"sync/atomic"

	"github.com/SunSetPilot/cs5296-project/client/config"
	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/utils"
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

	svcCtx := &ServiceContext{}
	svcCtx.SvcConf = c
	svcCtx.ClientStatus = atomic.Uint32{}
	svcCtx.ClientStatus.Store(model.CLIENT_STATUS_FREE)
	svcCtx.PodName = "mock-pod-name"
	svcCtx.PodUID = "mock-pod-uid"
	svcCtx.PodIP = "127.0.0.1"
	svcCtx.NodeName = "mock-node-name"
	svcCtx.NodeIP = "127.0.0.1"
	svcCtx.ServerAddr = "http://server.mock:8080"

	utils.MockHttpClient()
	return svcCtx
}
