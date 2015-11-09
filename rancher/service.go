package rancher

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher/client"
)

//ScaleService increases the scale of named service.
func (c *Context) ScaleService(serviceName string) error {
	cowbellService, err := c.getService(serviceName)
	if err != nil {
		return err
	}
	rancherService, err := c.FindExisting(serviceName)
	if err != nil {
		return err
	}

	newScale := rancherService.Scale + cowbellService.Increment

	if rancherService == nil {
		return fmt.Errorf("Failed to find %s to scale", serviceName)
	}

	logrus.Debugf("Setting %s scale to %d", serviceName, newScale)
	rancherService, err = c.Client.Service.Update(rancherService, map[string]interface{}{
		"scale": newScale,
	})
	if err != nil {
		return err
	}

	return c.Wait(rancherService)
}

//FindExisting retrieves an existing service in a stack
func (c *Context) FindExisting(name string) (*client.Service, error) {
	logrus.Debugf("Finding service %s", name)

	name, environmentID, err := c.resolveServiceAndEnvironmentID(name)
	if err != nil {
		return nil, err
	}

	services, err := c.Client.Service.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"environmentID": environmentID,
			"name":          name,
			"removed_null":  nil,
		},
	})

	if err != nil {
		return nil, err
	}

	if len(services.Data) == 0 {
		return nil, nil
	}

	logrus.Debugf("Found service %s", name)
	return &services.Data[0], nil
}

func (c *Context) resolveServiceAndEnvironmentID(name string) (string, string, error) {
	parts := strings.SplitN(name, "/", 2)
	if len(parts) == 1 {
		return name, c.stack.Id, nil
	}

	envs, err := c.Client.Environment.List(&client.ListOpts{
		Filters: map[string]interface{}{
			"name":         parts[0],
			"removed_null": nil,
		},
	})

	if err != nil {
		return "", "", err
	}

	if len(envs.Data) == 0 {
		return "", "", fmt.Errorf("Failed to find stack: %s", parts[0])
	}

	return parts[1], envs.Data[0].Id, nil
}
