package job

import (
	"github.com/SunSetPilot/cs5296-project/client/svc"
)

var Jobs []Job

type Job interface {
	GetName() string
	Do(ctx *svc.ServiceContext)
}
