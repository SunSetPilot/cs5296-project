package job

import (
	"encoding/json"
	"time"

	"cs5296-project/client/svc"
	"cs5296-project/model/request"
	"cs5296-project/utils"
	"cs5296-project/utils/log"
)

func init() {
	Jobs = append(Jobs, &SendHeartbeatJob{})
}

type SendHeartbeatJob struct {
}

func (j *SendHeartbeatJob) GetName() string {
	return "send_heartbeat_job"
}

func (j *SendHeartbeatJob) Do(ctx *svc.ServiceContext) {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("%s panic: %v", j.GetName(), r)
		}
	}()
	ticker := time.NewTicker(time.Duration(ctx.SvcConf.HeartbeatInterval) * time.Second)
	for {
		select {
		case <-ticker.C:
			req := &request.HeartbeatRequest{
				PodName:      ctx.PodName,
				PodUID:       ctx.PodUID,
				PodIP:        ctx.PodIP,
				NodeName:     ctx.NodeName,
				NodeIP:       ctx.NodeIP,
				ClientStatus: uint8(ctx.ClientStatus.Load()),
			}
			data, err := json.Marshal(req)
			if err != nil {
				log.Errorf("send_heartbeat_job json marshal failed: %v", err)
			}
			_, err = utils.HttpRequest(
				"POST",
				ctx.ServerAddr+"/api/v1/internal/heartbeat",
				string(data),
				nil,
				nil,
				false,
			)
			if err != nil {
				log.Errorf("send_heartbeat_job http request failed: %v", err)
			}
		}
	}
}
