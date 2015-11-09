package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// ScaleUp API entry point to scale service up
func ScaleUp(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	response := respondUnauthorized()

	logrus.Infof("Request to scale service: %s", variables["serviceName"])
	if checkServiceToken(variables["serviceName"], r.URL.Query()["token"][0]) {
		logrus.Infof("Scaling service...%s", variables["serviceName"])
		response = respondOK()
		if err := context.ScaleService(variables["serviceName"]); err != nil {
			logrus.Errorf("%s", err)
			response = respondServerError()
		}
	}

	output, _ := json.Marshal(response)
	fmt.Fprintln(w, string(output))
}

func respondServerError() *Response {
	return &Response{
		Type:   "error",
		Status: 500,
		Code:   http.StatusText(500),
	}
}
func respondUnauthorized() *Response {
	return &Response{
		Type:   "error",
		Status: 401,
		Code:   http.StatusText(401),
	}
}

func respondOK() *Response {
	return &Response{
		Type:   "message",
		Status: 200,
		Code:   http.StatusText(200),
	}
}

func checkServiceToken(serviceName string, token string) bool {
	match := false
	if token == context.GetServiceToken(serviceName) {
		match = true
	}
	return match
}
