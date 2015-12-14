package server

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/cloudnautique/cowbell/app"
)

var (
	context *app.Context
)

//StartServer starts the API server
func StartServer() {
	logrus.SetLevel(logrus.DebugLevel)
	context = &app.Context{}
	err := context.InitConfig()
	if err != nil {
		logrus.Fatalf("Could not load context: %s", err)
	}

	v1Router := NewRouter()
	logrus.Fatal(http.ListenAndServe(":8088", v1Router))
}
