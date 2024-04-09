package external

import (
	"fmt"
	"net/http"

	"cs5296-project/server/table"
	"cs5296-project/utils"
	"cs5296-project/utils/log"

	"github.com/gin-gonic/gin"
)

func (l *Logic) GetClientList(c *gin.Context) {
	var (
		rsp     *utils.Rsp
		clients []*table.TableClientModel
		err     error
	)
	rsp = utils.NewRsp(c)

	clients, err = table.TableClient.GetOnlineClientList(c.Request.Context())
	if err != nil {
		log.Errorf("GetClientList failed to get clients: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get clients"))
	}
	rsp.RspSuccess(clients)
}
