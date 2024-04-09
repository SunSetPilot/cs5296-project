package external

import (
	"cs5296-project/server/model"
	"cs5296-project/server/utils"
	"cs5296-project/server/utils/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (l *Logic) GetClientList(c *gin.Context) {
	var (
		rsp     *utils.Rsp
		clients []*model.TableClientModel
		err     error
	)
	rsp = utils.NewRsp(c)

	clients, err = model.TableClient.GetOnlineClientList(c.Request.Context())
	if err != nil {
		log.Errorf("GetClientList failed to get clients: %v", err)
		rsp.RspError(http.StatusInternalServerError, fmt.Errorf("failed to get clients"))
	}
	rsp.RspSuccess(clients)
}
