package job

import "cs5296-project/server/svc"

var Jobs []Job

type Job interface {
	GetName() string
	Do(ctx *svc.ServiceContext)
}
