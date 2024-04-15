package external

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/model/request"
	"github.com/SunSetPilot/cs5296-project/model/table"
	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/utils"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

func (l *Logic) CreateTask(c *gin.Context) {
	var (
		req []request.CreateTaskRequest
		rsp *utils.Rsp
		err error
	)
	rsp = utils.NewRsp(c)

	err = c.BindJSON(&req)
	log.Infof("CreateTask request: %v", req)
	if err != nil {
		log.Errorf("CreateTask failed to bind request: %v", err)
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}

	taskModels := make([]*table.TaskModel, 0)
	for _, task := range req {
		taskModels = append(taskModels, &table.TaskModel{
			TaskID:     uuid.NewString(),
			SrcPodIP:   task.SrcPodIP,
			SrcPodUID:  task.SrcPodUID,
			DstPodIP:   task.DstPodIP,
			DstPodUID:  task.DstPodUID,
			TaskParam:  task.TaskParam,
			TaskType:   task.TaskType,
			TaskStatus: model.TASK_STATUS_CREATED,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		})

	}

	err = dal.TableTask.BatchCreate(c.Request.Context(), taskModels)
	if err != nil {
		log.Errorf("CreateTask failed to batch create task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to create task"))
		return
	}
	result := make([]string, 0)
	for _, task := range taskModels {
		result = append(result, task.TaskID)
	}
	rsp.RspSuccess(result)
}
