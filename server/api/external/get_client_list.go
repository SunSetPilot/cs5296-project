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

func (l *Logic) GetClientList(c *gin.Context) {
	var (
		rsp     *utils.Rsp
		clients []*table.ClientModel
		err     error
	)
	rsp = utils.NewRsp(c)

	clients, err = dal.TableClient.GetOnlineClientList(c.Request.Context())
	if err != nil {
		log.Errorf("GetClientList failed to get clients: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get clients"))
	}
	rsp.RspSuccess(clients)
}
