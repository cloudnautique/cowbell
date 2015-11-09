package rancher

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher/client"
)

//LoadStack gets the rancher stack configuration.
func (c *Context) LoadStack() (*client.Environment, error) {
	if c.stack != nil {
		return c.stack, nil
	}

	projectName := getStackName()
	if err := c.loadRancherClient(); err != nil {
		return nil, err
	}

	logrus.Debugf("Looking for stack %s", projectName)
	// First try by name
	envs, err := c.Client.Environment.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"name":         projectName,
			"removed_null": nil,
		},
	})
	if err != nil {
		return nil, err
	}

	for _, env := range envs.Data {
		if strings.EqualFold(projectName, env.Name) {
			logrus.Debugf("Found stack: %s(%s)", env.Name, env.Id)
			c.stack = &env
			return c.stack, nil
		}
	}

	// Now try not by name for case sensitive databases
	envs, err = c.Client.Environment.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"removed_null": nil,
		},
	})
	if err != nil {
		return nil, err
	}

	for _, env := range envs.Data {
		if strings.EqualFold(projectName, env.Name) {
			logrus.Debugf("Found stack: %s(%s)", env.Name, env.Id)
			c.stack = &env
			return c.stack, nil
		}
	}

	logrus.Infof("Creating stack %s", projectName)
	env, err := c.Client.Environment.Create(&client.Environment{
		Name: projectName,
	})
	if err != nil {
		return nil, err
	}

	c.stack = env

	return c.stack, nil
}
