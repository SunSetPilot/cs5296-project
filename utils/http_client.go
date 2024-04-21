package utils

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/jarcoal/httpmock"

	"github.com/SunSetPilot/cs5296-project/model"
	"github.com/SunSetPilot/cs5296-project/model/table"
)

var (
	client *resty.Client
)

func init() {
	client = resty.New().SetTimeout(2 * time.Second)
}

func HttpRequest(method, url, data string, params, headers map[string]string, json bool) (map[string]interface{}, error) {
	var (
		err      error
		response *resty.Response
		request  *resty.Request
		resp     map[string]interface{}
	)
	resp = make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	request = client.R()
	request.Method = method
	request.URL = url

	if json {
		_, err = request.SetQueryParams(params).SetHeaders(headers).SetBody(data).SetResult(&resp).Send()
	} else {
		response, err = request.SetQueryParams(params).SetHeaders(headers).SetBody(data).Send()
		resp["url"] = response.Request.URL
		resp["status_code"] = response.StatusCode()
		resp["headers"] = response.Header()
		resp["body"] = response.String()
	}

	return resp, err
}

func MockHttpClient() {
	httpmock.ActivateNonDefault(client.GetClient())

	// mock heartbeat response
	heartbeatResponder, _ := httpmock.NewJsonResponder(200, Rsp{
		Status: 0,
		Msg:    "",
		Data:   nil,
	})
	httpmock.RegisterResponder(
		"POST",
		"http://server.mock:8080/api/v1/internal/heartbeat",
		heartbeatResponder,
	)

	// mock report_task response
	reportTaskResponder, _ := httpmock.NewJsonResponder(200, Rsp{
		Status: 0,
		Msg:    "",
		Data:   nil,
	})
	httpmock.RegisterResponder(
		"POST",
		"http://server.mock:8080/api/v1/internal/report_task",
		reportTaskResponder,
	)

	// mock fetch_tasks response
	fetchTasksResponder, _ := httpmock.NewJsonResponder(200, []*table.TaskModel{
		{
			ID:         1,
			TaskID:     uuid.NewString(),
			SrcPodUID:  "src-pod-uid",
			SrcPodIP:   "127.0.0.1",
			DstPodUID:  "dst-pod-uid",
			DstPodIP:   "127.0.0.1",
			TaskParam:  "task-param",
			TaskType:   "ping",
			TaskStatus: model.TASK_STATUS_CREATED,
			TaskResult: "",
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	})
	httpmock.RegisterResponder(
		"GET",
		"http://server.mock:8080/api/v1/internal/get_tasks",
		fetchTasksResponder,
	)
}
