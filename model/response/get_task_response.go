package response

type Task struct {
	TaskID    string
	TaskType  string
	DstPodIP  string
	TaskParam string
}

type GetTaskResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []Task `json:"data"`
}
