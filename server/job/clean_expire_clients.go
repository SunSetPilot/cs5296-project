package job

import (
	"context"
	"cs5296-project/server/model"
	"cs5296-project/server/svc"
	"cs5296-project/server/utils/log"
	"time"
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
	timer := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-timer.C:
			err = model.TableClient.UpdateOfflineClients(context.Background())
			if err != nil {
				log.Errorf("clean_expire_clients_job error: %v", err)
			}
		}
	}
}
