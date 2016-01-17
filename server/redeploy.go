package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// ScaleUp API entry point to scale service up
func ReDeployEntry(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	response := respond("message", 404)

	if tokenList, ok := r.URL.Query()["token"]; ok {
		logrus.Infof("Request to redeploy service: %s", variables["serviceName"])
		if checkServiceToken(variables["serviceName"], tokenList[0]) {
			logrus.Infof("Redeploy!")
			context.Redeploy(variables["serviceName"])
		}
	}

	output, _ := json.Marshal(response)
	fmt.Fprintln(w, string(output))
}
