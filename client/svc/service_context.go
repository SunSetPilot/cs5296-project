package svc

import (
	"fmt"
	"os"
	"sync/atomic"

	"github.com/SunSetPilot/cs5296-project/client/config"
	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

type ServiceContext struct {
	SvcConf *config.Config

	ClientStatus atomic.Uint32

	PodName    string
	PodUID     string
	PodIP      string
	NodeName   string
	NodeIP     string
	ServerAddr string
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

	svcCtx := &ServiceContext{}
	svcCtx.SvcConf = c
	svcCtx.ClientStatus = atomic.Uint32{}
	svcCtx.ClientStatus.Store(model.CLIENT_STATUS_FREE)
	svcCtx.PodName = os.Getenv("POD_NAME")
	if svcCtx.PodName == "" {
		panic("POD_NAME env is empty")
	}
	svcCtx.PodUID = os.Getenv("POD_UID")
	if svcCtx.PodUID == "" {
		panic("POD_UID env is empty")
	}
	svcCtx.PodIP = os.Getenv("POD_IP")
	if svcCtx.PodIP == "" {
		panic("POD_IP env is empty")
	}
	svcCtx.NodeName = os.Getenv("NODE_NAME")
	if svcCtx.NodeName == "" {
		panic("NODE_NAME env is empty")
	}
	svcCtx.NodeIP = os.Getenv("NODE_IP")
	if svcCtx.NodeIP == "" {
		panic("NODE_IP env is empty")
	}
	svcCtx.ServerAddr = os.Getenv("SERVER_ADDR")
	if svcCtx.ServerAddr == "" {
		panic("SERVER_ADDR env is empty")
	}

	return svcCtx
}
