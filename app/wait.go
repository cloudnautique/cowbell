package app

import (
	"time"

	"github.com/rancher/go-rancher/client"
)

// WaitFor waits for a resource to reach a certain state.
func (c *Context) WaitFor(resource *client.Resource, output interface{}, transitioning func() string) error {
	for {
		if transitioning() != "yes" {
			return nil
		}

		time.Sleep(150 * time.Millisecond)

		err := c.Client.Reload(resource, output)
		if err != nil {
			return err
		}
	}
}

// Wait wait for a service resource to transition
func (c *Context) Wait(service *client.Service) error {
	return c.WaitFor(&service.Resource, service, func() string {
		return service.Transitioning
	})
}
