package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"cs5296-project/client/config"
	"cs5296-project/client/job"
	"cs5296-project/client/svc"
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

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)
	select {
	case sig := <-ch:
		fmt.Printf("Received signal %s, exiting...\n", sig)
		os.Exit(0)
	}
}
