package internal

import (
	"github.com/SunSetPilot/cs5296-project/server/svc"
)

type Logic struct {
	svcCtx *svc.ServiceContext
}

func NewInternalLogic(svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		svcCtx: svcCtx,
	}
}
