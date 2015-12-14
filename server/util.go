package server

import (
	"net/http"
)

func respond(responseType string, code int) *Response {
	return &Response{
		Type:   responseType,
		Status: code,
		Code:   http.StatusText(code),
	}
}

func checkServiceToken(serviceName string, token string) bool {
	match := false
	if token == context.GetServiceToken(serviceName) {
		match = true
	}
	return match
}
