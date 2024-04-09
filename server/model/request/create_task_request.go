package request

type CreateTaskRequest struct {
	SrcPodUID string `json:"src_pod_uid"`
	SrcPodIP  string `json:"src_pod_ip"`
	DstPodUID string `json:"dst_pod_uid"`
	DstPodIP  string `json:"dst_pod_ip"`
	TaskParam string `json:"task_param"`
	TaskType  string `json:"task_type"`
}
