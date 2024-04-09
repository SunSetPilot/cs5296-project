package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	client     *resty.Client
	clientOnce sync.Once
)

func getClient() *resty.Client {
	clientOnce.Do(func() {
		client = resty.New().SetTimeout(2 * time.Second)
	})
	return client
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

	request = getClient().R()
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
