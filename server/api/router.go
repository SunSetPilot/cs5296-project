package api

import (
	"github.com/gin-gonic/gin"

	"cs5296-project/server/api/external"
	"cs5296-project/server/api/internal"
	"cs5296-project/server/svc"
)

func RegisterRoutes(server *gin.Engine, ctx *svc.ServiceContext) {
	server.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
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
