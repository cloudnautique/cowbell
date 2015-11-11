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
	response := respond("message", 404)

	logrus.Infof("Request to scale service: %s", variables["serviceName"])
	if checkServiceToken(variables["serviceName"], r.URL.Query()["token"][0]) {
		logrus.Infof("Scaling service...%s", variables["serviceName"])

		//quickly grab the token and respond. This action can take a while.
		go context.ScaleServiceUp(variables["serviceName"])
		response = respond("message", 202)
	}

	output, _ := json.Marshal(response)
	fmt.Fprintln(w, string(output))
}

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
