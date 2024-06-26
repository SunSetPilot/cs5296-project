package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SunSetPilot/cs5296-project/server/api/external"
	"github.com/SunSetPilot/cs5296-project/server/api/internal"
	"github.com/SunSetPilot/cs5296-project/server/svc"
)

func RegisterRoutes(server *gin.Engine, ctx *svc.ServiceContext) {
	server.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	internalRoutes := server.Group("/api/v1/internal")
	internalLogic := internal.NewInternalLogic(ctx)
	internalRoutes.POST("/heartbeat", internalLogic.HeartBeat)
	internalRoutes.GET("/get_tasks", internalLogic.GetTasks)
	internalRoutes.POST("/report_task", internalLogic.ReportTask)

	externalRoutes := server.Group("/api/v1/external")
	externalLogic := external.NewExternalLogic(ctx)
	externalRoutes.GET("/clients", externalLogic.GetClientList)
	externalRoutes.POST("/task/create", externalLogic.CreateTask)
	externalRoutes.GET("/task", externalLogic.GetTask)
}
