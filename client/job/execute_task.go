package job

import (
	"github.com/SunSetPilot/cs5296-project/client/pool"
	"github.com/SunSetPilot/cs5296-project/client/svc"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

func init() {
	Jobs = append(Jobs, &ExecuteTaskJob{})
}

type ExecuteTaskJob struct {
}

func (j *ExecuteTaskJob) GetName() string {
	return "execute_task_job"
}

func (j *ExecuteTaskJob) Do(ctx *svc.ServiceContext) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("%s panic: %v", j.GetName(), r)
		}
	}()
	execPool := pool.NewCommandExecPool(ctx.SvcConf.ExecPoolSize)
	execPool.Start()

}
