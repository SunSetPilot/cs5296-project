package external

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SunSetPilot/cs5296-project/model/table"
	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/utils"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

func (l *Logic) GetTask(c *gin.Context) {
	var (
		rsp  *utils.Rsp
		task *table.TaskModel
		err  error
	)
	rsp = utils.NewRsp(c)

	taskID := c.Query("task_id")
	log.Infof("GetTask request: %v", taskID)
	if taskID == "" {
		log.Errorf("task_id is required")
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("task_id is required"))
		return
	}

	task, err = dal.TableTask.GetTaskByTaskID(c.Request.Context(), taskID)
	if err != nil {
		log.Errorf("GetTask failed to get task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get task"))
		return
	}
	rsp.RspSuccess(task)
}
