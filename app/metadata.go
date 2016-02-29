package app

import (
	"github.com/Sirupsen/logrus"
	"github.com/rancher/go-rancher-metadata/metadata"
)

const (
	metadataURL = "http://rancher-metadata"
)

func getServiceMetadata() map[string]interface{} {
	serviceMetadata := map[string]interface{}{}
	md, err := metadata.NewClientAndWait(metadataURL + "/2015-12-19")
	if err != nil {
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
