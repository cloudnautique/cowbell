package app

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher/client"
)

// Context contains the application config context
type Context struct {
	config     *cowbellConfig
	Client     *client.RancherClient
	stack      *client.Environment
	RancherURL string
	accessKey  string
	secretKey  string
}

type cowbellConfig struct {
	services map[string]service
}

type service struct {
	Name             string   `json:"name,omitempty"`
	Decrement        int64    `json:"decrement,omitempty"`
	Increment        int64    `json:"increment,omitempty"`
	Token            string   `json:"token,omitempty"`
	QuietTime        int      `json:"quietTime,omitempty"`
	quietTimeChannel chan int `json:"-"`
	PullOnUpgrade    bool     `json:"pullOnUpgrade,omitempty"`
}

//InitConfig initializes the application context.
func (c *Context) InitConfig() error {
	if err := c.loadConfigFromMetadata(); err != nil {
		return err
	}

	c.loadRancherClient()
	c.LoadStack()

	return nil
}

//GetServiceToken gets the configured service token from rancher metadata.
func (c *Context) GetServiceToken(serviceName string) string {
	token := ""
	if serviceConfig, ok := c.config.services[serviceName]; ok {
		token = serviceConfig.Token
	}
	return token
}

func (c *Context) getService(serviceName string) (*service, error) {
	cowbellService, ok := c.config.services[serviceName]
	if !ok {
		return &service{}, fmt.Errorf("Could not load service config for %s", serviceName)
	}

	return &cowbellService, nil

}

func (c *Context) loadConfigFromMetadata() error {
	c.config = &cowbellConfig{
		services: map[string]service{},
	}

	metadata := getServiceMetadata()

	if c.RancherURL = os.Getenv("CATTLE_URL"); c.RancherURL == "" {
		logrus.Fatalf("Could not find Rancher URL")
	}

	if c.accessKey = os.Getenv("CATTLE_ACCESS_KEY"); c.accessKey == "" {
		logrus.Fatalf("Could not find Rancher access key")
	}

	if c.secretKey = os.Getenv("CATTLE_SECRET_KEY"); c.secretKey == "" {
		logrus.Fatalf("Could not find Rancher secret key")
	}

	if _, ok := metadata["services"]; ok {
		for _, service := range metadata["services"].([]interface{}) {
			s := setService(service.(map[string]interface{}))
			logrus.Infof("Found Cowbell config for service: %s", s.Name)
			c.config.services[s.Name] = *s
		}
	}
	return nil
}

func setService(serviceDef map[string]interface{}) *service {
	s := &service{}
	bytesData, err := convertBytesFromMapStringInterface(serviceDef)
	if err != nil {
		logrus.Error(err)
		logrus.Fatalf("Can't load service configs")
	}
	err = json.Unmarshal(bytesData, &s)
	if err != nil {
		logrus.Error(err)
		logrus.Fatalf("Can't load service configs")
	}

	s.quietTimeChannel = make(chan int, 1)

	return s
}

func convertBytesFromMapStringInterface(data map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func getIntFromFloat64(n interface{}) int {
	return int(n.(float64))
}

func getInt64FromFloat64(n interface{}) int64 {
	return int64(n.(float64))
}

func (c *Context) loadRancherClient() error {
	if c.Client == nil {
		if c.RancherURL == "" {
			return fmt.Errorf("RancherURL is nil, can not change Scale")
		}

		client, err := client.NewRancherClient(&client.ClientOpts{
			Url:       c.RancherURL,
			AccessKey: c.accessKey,
			SecretKey: c.secretKey,
		})
		if err != nil {
			return err
		}

		c.Client = client
	}
	return nil
}
