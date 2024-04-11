package external

import (
	"github.com/SunSetPilot/cs5296-project/server/svc"
)

type Logic struct {
	svcCtx *svc.ServiceContext
}

func NewExternalLogic(svcCtx *svc.ServiceContext) *Logic {
	return &Logic{
		svcCtx: svcCtx,
	}
}
