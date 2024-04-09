package request

type HeartbeatRequest struct {
	PodName      string `json:"pod_name"`
	PodUID       string `json:"pod_uid"`
	PodIP        string `json:"pod_ip"`
	NodeName     string `json:"node_name"`
	NodeIP       string `json:"node_ip"`
	ClientStatus uint8  `json:"client_status"`
}
