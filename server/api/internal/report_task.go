package internal

import (
	"fmt"
	"net/http"
	"time"

	"cs5296-project/model/request"
	"cs5296-project/server/table"
	"cs5296-project/utils"
	"cs5296-project/utils/log"

	"github.com/gin-gonic/gin"
)

func (l *Logic) ReportTask(c *gin.Context) {
	var (
		req request.ReportTaskRequest
		rsp *utils.Rsp
		err error
	)
	rsp = utils.NewRsp(c)

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("ReportTask failed to bind request: %v", err)
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}

	taskModel := &table.TableTaskModel{
		TaskID:     req.TaskID,
		TaskStatus: req.TaskStatus,
		TaskResult: req.TaskResult,
		UpdateTime: time.Now(),
	}

	err = table.TableTask.UpdateByTaskID(c.Request.Context(), taskModel)
	if err != nil {
		log.Errorf("ReportTask failed to update task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to update task"))
		return
	}
	rsp.RspSuccess(nil)
}
