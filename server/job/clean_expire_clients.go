package job

import (
	"context"
	"time"

	"github.com/SunSetPilot/cs5296-project/server/dal"
	"github.com/SunSetPilot/cs5296-project/server/svc"
	"github.com/SunSetPilot/cs5296-project/utils/log"
)

func init() {
	Jobs = append(Jobs, &CleanExpireClientsJob{})
}

type CleanExpireClientsJob struct {
}

func (j *CleanExpireClientsJob) GetName() string {
	return "clean_expire_clients_job"
}

func (j *CleanExpireClientsJob) Do(ctx *svc.ServiceContext) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("clean_expire_clients_job panic: %v", r)
		}
	}()
	var err error

	for {
		time.Sleep(3 * time.Second)
		log.Debugf("clean_expire_clients_job running")
		err = dal.TableClient.UpdateOfflineClients(context.Background(), ctx.SvcConf.ClientOfflineThreshold)
		if err != nil {
			log.Errorf("clean_expire_clients_job error: %v", err)
		}
		log.Debugf("clean_expire_clients_job done")
	}
}
