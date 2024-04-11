package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/SunSetPilot/cs5296-project/server/api"
	"github.com/SunSetPilot/cs5296-project/server/config"
	"github.com/SunSetPilot/cs5296-project/server/job"
	"github.com/SunSetPilot/cs5296-project/server/svc"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	config.MustLoad(*configFile, &c)

	ctx := svc.MustNewServiceContext(&c)

	for _, j := range job.Jobs {
		go j.Do(ctx)
		fmt.Printf("Job %s started\n", j.GetName())
	}

	server := gin.Default()
	api.RegisterRoutes(server, ctx)

	fmt.Printf("Server is running on %s\n", c.ListenOn)
	if err := server.Run(c.ListenOn); err != nil {
		panic(fmt.Errorf("failed to start server: %w", err))
	}

}
