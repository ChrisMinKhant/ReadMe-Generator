package generatereadmefileservice

import (
	"strings"

	dto "github.com/ChrisMinKhant/megoyougo_framework/dto/request"
	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"github.com/sirupsen/logrus"
)

type generateReadmeFileServiceImplementation struct {
}

func New() *generateReadmeFileServiceImplementation {
	return &generateReadmeFileServiceImplementation{}
}

func (generateReadmeFileServiceImplementation *generateReadmeFileServiceImplementation) GenerateReadmeFile(serviceInfo *dto.GenerateReadmeFileRequest) {
	logrus.Infof("Fetched service info ::: %v\n", serviceInfo)

	header := util.HEADER

	// Set value to header of the readmefile
	generateReadmeFileServiceImplementation.setTemplateValue("[ServiceName]", serviceInfo.ServiceName, &header)
	generateReadmeFileServiceImplementation.setTemplateValue("[ServiceCategory]", serviceInfo.ServiceCategory, &header)
	generateReadmeFileServiceImplementation.setTemplateValue("[ServiceDescription]", serviceInfo.Description, &header)

	body := ""

	// Set value to body of the readmefile
	for index, endpoint := range serviceInfo.Endpoints {
		tempBody := util.Endpoint

		generateReadmeFileServiceImplementation.setTemplateValue("[EndpointNumber]", string(index+1), &tempBody)
		generateReadmeFileServiceImplementation.setTemplateValue("[Path]", endpoint.Path, &tempBody)
		generateReadmeFileServiceImplementation.setTemplateValue("[ExampleRequest]", endpoint.ExampleRequest, &tempBody)
		generateReadmeFileServiceImplementation.setTemplateValue("[ExampleResponse]", endpoint.ExampleResponse, &tempBody)

		body += tempBody
	}

	logrus.Infof("Fetched full readmefile ::: %v\n", header+body)
}

func (generateReadmeFileServiceImplementation *generateReadmeFileServiceImplementation) setTemplateValue(key string, value string, template *string) {
	*template = strings.Replace(*template, key, value, 1)
}
