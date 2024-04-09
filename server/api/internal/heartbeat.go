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

func (l *Logic) HeartBeat(c *gin.Context) {
	var (
		req request.HeartbeatRequest
		rsp *utils.Rsp
		err error
	)
	rsp = utils.NewRsp(c)

	err = c.BindJSON(&req)
	if err != nil {
		log.Errorf("HeartBeat failed to bind request: %v", err)
		rsp.RspError(http.StatusBadRequest, fmt.Errorf("invalid request"))
		return
	}

	clientModel := &table.TableClientModel{
		PodName:      req.PodName,
		PodUID:       req.PodUID,
		PodIP:        req.PodIP,
		NodeName:     req.NodeName,
		NodeIP:       req.NodeIP,
		ClientStatus: req.ClientStatus,
		RegisterTime: time.Now(),
		UpdateTime:   time.Now(),
	}

	err = table.TableClient.CreateOrUpdate(c.Request.Context(), clientModel)
	if err != nil {
		log.Errorf("HeartBeat failed to create or update client: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to create or update client"))
		return
	}
	rsp.RspSuccess(nil)
}
