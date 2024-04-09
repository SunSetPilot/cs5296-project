package request

type ReportTaskRequest struct {
	TaskID     string `json:"task_id"`
	TaskStatus uint8  `json:"task_status"`
	TaskResult string `json:"task_result"`
}
