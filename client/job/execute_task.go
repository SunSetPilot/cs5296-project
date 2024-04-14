package job

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SunSetPilot/cs5296-project/client/pool"
	"github.com/SunSetPilot/cs5296-project/client/svc"
	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/model/request"
	"github.com/SunSetPilot/cs5296-project/model/response"
	"github.com/SunSetPilot/cs5296-project/utils"
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

	for {
		select {
		case result := <-execPool.Out:
			if result.Err != nil {
				log.Errorf("execute_task_job command execute failed: %v, task_id: %s", result.Err, result.ID)
				err := reportTask(ctx, &request.ReportTaskRequest{
					TaskID:     result.ID,
					TaskStatus: model.TASK_STATUS_FAILED,
					TaskResult: result.Err.Error(),
				})
				if err != nil {
					log.Errorf("execute_task_job report task failed: %v", err)
				}
			} else {
				err := reportTask(ctx, &request.ReportTaskRequest{
					TaskID:     result.ID,
					TaskStatus: model.TASK_STATUS_FINISHED,
					TaskResult: string(result.Result),
				})
				if err != nil {
					log.Errorf("execute_task_job report task failed: %v", err)
				}
			}
		default:
			time.Sleep(time.Second)
			tasks, err := fetchTasks(ctx)
			if err != nil {
				log.Errorf("execute_task_job fetch tasks failed: %v", err)
				continue
			}
			for _, task := range tasks {
				if execPool.IsFull() {
					break
				}
				var destArg string
				if task.TaskType == model.TASK_TYPE_PING || task.TaskType == model.TASK_TYPE_TRACEROUTE {
					destArg = task.DstPodIP
				} else if task.TaskType == model.TASK_TYPE_IPERF {
					destArg = "-c " + task.DstPodIP
				} else {
					log.Errorf("execute_task_job invalid task type: %s", task.TaskType)
					continue
				}
				execPool.In <- &pool.Command{
					ID:   task.TaskID,
					Name: task.TaskType,
					Args: []string{destArg, task.TaskParam},
				}
				err = reportTask(ctx, &request.ReportTaskRequest{
					TaskID:     task.TaskID,
					TaskStatus: model.TASK_STATUS_RUNNING,
				})
				if err != nil {
					log.Errorf("execute_task_job report task failed: %v", err)
				}
			}
		}
	}

}

func fetchTasks(ctx *svc.ServiceContext) ([]response.Task, error) {
	var (
		resp  map[string]interface{}
		tasks response.GetTaskResponse
		err   error
	)

	resp, err = utils.HttpRequest(
		"GET",
		ctx.ServerAddr+"/api/v1/internal/get_tasks",
		"",
		map[string]string{
			"pod_uid": ctx.PodUID,
		},
		nil,
		false,
	)
	if err != nil {
		log.Errorf("http request failed: %v", err)
		return nil, err
	}
	body := resp["body"].(string)
	err = json.Unmarshal([]byte(body), &tasks)
	if err != nil {
		log.Errorf("json unmarshal failed: %v", err)
		return nil, err
	}
	return tasks.Data, nil
}

func reportTask(ctx *svc.ServiceContext, req *request.ReportTaskRequest) error {
	var (
		err  error
		data []byte
	)
	data, err = json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshal failed: %v", err)
	}
	_, err = utils.HttpRequest(
		"POST",
		ctx.ServerAddr+"/api/v1/internal/report_task",
		string(data),
		nil,
		nil,
		false,
	)
	if err != nil {
		return fmt.Errorf("http request failed: %v", err)
	}
	return nil
}
