package rancher

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher-metadata/metadata"
)

const (
	metadataURL = "http://rancher-metadata"
)

func getServiceMetadata() map[string]interface{} {
	serviceMetadata := map[string]interface{}{}
	md := metadata.NewClient(metadataURL + "/2015-07-25")
	if err := testConnection(md); err != nil {
		logrus.Fatalf("Can not load configuration from metadata")
	}

	serviceData, err := md.GetSelfService()
	if err != nil {
		logrus.Errorf("%s", err)
	}

	serviceMetadata = serviceData.Metadata

	return serviceMetadata
}

func getStackName() string {
	md := metadata.NewClient(metadataURL + "/2015-07-25")

	stackData, err := md.GetSelfStack()
	if err != nil {
		logrus.Errorf("Could not get Stack Name")
		return ""
	}

	return stackData.Name

}

func testConnection(mdClient *metadata.Client) error {
	var err error
	maxTime := 20 * time.Second

	for i := 1 * time.Second; i < maxTime; i *= time.Duration(2) {
		if _, err = mdClient.GetVersion(); err != nil {
			time.Sleep(i)
		} else {
			return nil
		}
	}
	return err
}
