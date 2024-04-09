package internal

import (
	"cs5296-project/server/model"
	"cs5296-project/server/model/request"
	"cs5296-project/server/utils"
	"cs5296-project/server/utils/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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

	taskModel := &model.TableTaskModel{
		TaskID:     req.TaskID,
		TaskStatus: req.TaskStatus,
		TaskResult: req.TaskResult,
		UpdateTime: time.Now(),
	}

	err = model.TableTask.UpdateByTaskID(c.Request.Context(), taskModel)
	if err != nil {
		log.Errorf("ReportTask failed to update task: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to update task"))
		return
	}
	rsp.RspSuccess(nil)
}
