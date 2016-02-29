package app

import (
	"github.com/Sirupsen/logrus"
)

func (c *Context) Redeploy(serviceName string) error {
	rancherService, _ := c.FindExisting(serviceName)
	logrus.Debug(rancherService.LaunchConfig)
	if rancherService.State != "active" && rancherService.State != "inactive" {
		logrus.Errorf("Service %s must be state=active or state=inactive to be upgraded", serviceName)
	}

	return nil
}
