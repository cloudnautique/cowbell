package server

import (
	"github.com/gorilla/mux"
)

//NewRouter returns a new application HTTP router.
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1-scale/services/{serviceName}", ScaleUp).Methods("POST")
	return router
}
