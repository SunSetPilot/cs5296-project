package job

import (
	"cs5296-project/client/svc"
	"cs5296-project/utils/log"
)

func init() {
	Jobs = append(Jobs, &ExecuteTaskJob{})
}

type Task struct {
	TaskID    string
	TargetIP  string
	TaskParam string
}

type ExecuteTaskJob struct {
	taskChan chan *Task
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

}
