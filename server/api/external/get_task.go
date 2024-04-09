package external

import (
	"fmt"
	"net/http"

	"cs5296-project/server/table"
	"cs5296-project/utils"
	"cs5296-project/utils/log"

	"github.com/gin-gonic/gin"
)

func (l *Logic) GetTask(c *gin.Context) {
	var (
		rsp  *utils.Rsp
		task *table.TableTaskModel
		err  error
	)
	rsp = utils.NewRsp(c)

	taskID := c.Query("task_id")
	if taskID == "" {
		log.Errorf("task_id is required")
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("task_id is required"))
		return
	}

	task, err = table.TableTask.GetTaskByTaskID(c.Request.Context(), taskID)
	if err != nil {
		log.Errorf("GetTask failed to get task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get task"))
		return
	}
	rsp.RspSuccess(task)
}
