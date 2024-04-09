package external

import (
	"cs5296-project/server/model"
	"cs5296-project/server/utils"
	"cs5296-project/server/utils/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (l *Logic) GetTask(c *gin.Context) {
	var (
		rsp  *utils.Rsp
		task *model.TableTaskModel
		err  error
	)
	rsp = utils.NewRsp(c)

	taskID := c.Query("task_id")
	if taskID == "" {
		log.Errorf("task_id is required")
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("task_id is required"))
		return
	}

	task, err = model.TableTask.GetTaskByTaskID(c.Request.Context(), taskID)
	if err != nil {
		log.Errorf("GetTask failed to get task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get task"))
		return
	}
	rsp.RspSuccess(task)
}
