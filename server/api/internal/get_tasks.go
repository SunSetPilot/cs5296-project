package internal

import (
	"fmt"
	"net/http"

	"cs5296-project/server/table"
	"cs5296-project/utils"
	"cs5296-project/utils/log"

	"github.com/gin-gonic/gin"
)

func (l *Logic) GetTasks(c *gin.Context) {
	var (
		rsp   *utils.Rsp
		tasks []*table.TableTaskModel
		err   error
	)
	rsp = utils.NewRsp(c)

	podUID := c.Query("pod_uid")
	if podUID == "" {
		log.Errorf("pod_uid is required")
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("pod_uid is required"))
		return
	}

	tasks, err = table.TableTask.GetTaskBySrcPodUID(c.Request.Context(), podUID)
	if err != nil {
		log.Errorf("GetTasks failed to get tasks: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get tasks"))
		return
	}
	rsp.RspSuccess(tasks)
}
