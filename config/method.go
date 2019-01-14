package config

import (
	"fmt"
	"go-interface/request"
)

// URL ...
func URL(api string) string {
	return baseURL + api
}

// RequestID return id of the request
func RequestID(resp *request.Response) string {
	id := resp.Response.Header.Get("X-Request-Id")
	return id
}

// ErrInfo return the response text and the request_id
func ErrInfo(resp *request.Response) (msg string) {
	text, err := resp.Text()
	if err != nil {
		return "Response: " + text + "; Request_Id: " + "; Err: " + err.Error()
	}
	rid := RequestID(resp)
	msg = "Response: " + text + "; Request_Id: " + rid + "; Err: nil"
	return msg
}

// AssertInfo return request Data and Errinfo() message
func AssertInfo(acutal interface{}, resp *request.Response) string {
	r := ErrInfo(resp)
	msg := fmt.Sprintf("Data: %s; %v", acutal, r)
	return "\033[31m" + msg + "\033[0m"
}
