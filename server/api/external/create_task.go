package external

import (
	"cs5296-project/server/model"
	"cs5296-project/server/model/request"
	"cs5296-project/server/utils"
	"cs5296-project/server/utils/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (l *Logic) CreateTask(c *gin.Context) {
	var (
		req []request.CreateTaskRequest
		rsp *utils.Rsp
		err error
	)
	rsp = utils.NewRsp(c)

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("CreateTask failed to bind request: %v", err)
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}

	taskModels := make([]*model.TableTaskModel, 0)
	for _, task := range req {
		taskModels = append(taskModels, &model.TableTaskModel{
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

	err = model.TableTask.BatchCreate(c.Request.Context(), taskModels)
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
