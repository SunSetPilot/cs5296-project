package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SunSetPilot/cs5296-project/model/table"
	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/utils"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

func (l *Logic) GetTasks(c *gin.Context) {
	var (
		rsp   *utils.Rsp
		tasks []*table.TaskModel
		err   error
	)
	rsp = utils.NewRsp(c)

	podUID := c.Query("pod_uid")
	log.Infof("GetTasks request: %v", podUID)
	if podUID == "" {
		log.Errorf("pod_uid is required")
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("pod_uid is required"))
		return
	}

	tasks, err = dal.TableTask.GetTaskBySrcPodUID(c.Request.Context(), podUID)
	if err != nil {
		log.Errorf("GetTasks failed to get tasks: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get tasks"))
		return
	}
	rsp.RspSuccess(tasks)
}
