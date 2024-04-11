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

	clientModel := &table.ClientModel{
		PodName:      req.PodName,
		PodUID:       req.PodUID,
		PodIP:        req.PodIP,
		NodeName:     req.NodeName,
		NodeIP:       req.NodeIP,
		ClientStatus: req.ClientStatus,
		RegisterTime: time.Now(),
		UpdateTime:   time.Now(),
	}

	err = dal.TableClient.CreateOrUpdate(c.Request.Context(), clientModel)
	if err != nil {
		log.Errorf("HeartBeat failed to create or update client: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to create or update client"))
		return
	}
	rsp.RspSuccess(nil)
}
