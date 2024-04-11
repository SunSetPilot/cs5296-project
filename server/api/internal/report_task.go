package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/SunSetPilot/cs5296-project/model/request"
	"github.com/SunSetPilot/cs5296-project/model/table"
	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/utils"
	"github.com/SunSetPilot/cs5296-project/utils/log"
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

	taskModel := &table.TaskModel{
		TaskID:     req.TaskID,
		TaskStatus: req.TaskStatus,
		TaskResult: req.TaskResult,
		UpdateTime: time.Now(),
	}

	err = dal.TableTask.UpdateByTaskID(c.Request.Context(), taskModel)
	if err != nil {
		log.Errorf("ReportTask failed to update task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to update task"))
		return
	}
	rsp.RspSuccess(nil)
}
