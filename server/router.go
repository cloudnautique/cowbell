package server

import (
	"github.com/gorilla/mux"
)

//NewRouter returns a new application HTTP router.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1-scale/services/{serviceName}", ScaleUp).Methods("POST")
	router.HandleFunc("/v1-redeploy/services/{serviceName}", ReDeployEntry).Methods("POST")
	return router
}
