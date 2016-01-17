package app

import (
	"github.com/Sirupsen/logrus"
)

func (c *Context) Redeploy(serviceName string) error {
	rancherService, _ := c.FindExisting(serviceName)
	logrus.Debug(rancherService.LaunchConfig)

	return nil
}
