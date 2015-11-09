package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/cloudnautique/cowbell/server"
)

func main() {
	logrus.Infof("Starting Cowbell Server")
	server.StartServer()
}
