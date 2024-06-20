package generatereadmefileservice

import (
	"os"
	"strings"

	dto "github.com/ChrisMinKhant/megoyougo_framework/dto/request"
	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"github.com/sirupsen/logrus"
)

type generateReadmeFileServiceImplementation struct {
	exceptionHandler exception.Exception
}

func New() *generateReadmeFileServiceImplementation {
	return &generateReadmeFileServiceImplementation{
		exceptionHandler: exception.GetGeneralExceptionInstance(),
	}
}

func (generateReadmeFileServiceImplementation *generateReadmeFileServiceImplementation) GenerateReadmeFile(serviceInfo *dto.GenerateReadmeFileRequest) bool {
	defer generateReadmeFileServiceImplementation.exceptionHandler.RecoverPanic()

	logrus.Info("Generating readme file...")

	createdFile, error := os.Create(serviceInfo.GenerationDirectory + "/README.md")

	if error != nil {
		logrus.Panicf("Error occurred at creating file ::: %v\n", error.Error())
	}

	defer createdFile.Close()

	_, error = createdFile.Write([]byte(generateReadmeFileServiceImplementation.buildReadMe(serviceInfo)))

	if error != nil {
		logrus.Panicf("Error occurred at writing to file ::: %v\n", error.Error())
	}

	return true
}

func (generateReadmeFileServiceImplementation *generateReadmeFileServiceImplementation) buildReadMe(serviceInfo *dto.GenerateReadmeFileRequest) string {
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

		for _, requestParam := range endpoint.RequestParams {

			requestParamsSection := util.REQUEST_PARAM

			generateReadmeFileServiceImplementation.setTemplateValue("[FIELD]", requestParam.Field, &requestParamsSection)
			generateReadmeFileServiceImplementation.setTemplateValue("[TYPE]", requestParam.Type, &requestParamsSection)
			generateReadmeFileServiceImplementation.setTemplateValue("[REQUEST_PARAM_DESCRIPTION]", requestParam.Description, &requestParamsSection)

			validationsSection := ""

			for _, validation := range requestParam.Validations {
				validationsSection += "| " + validation.Name + " | " + validation.Description + " |\n"
			}

			generateReadmeFileServiceImplementation.setTemplateValue("[VALIDATIONS]", validationsSection, &requestParamsSection)

			tempBody += requestParamsSection
		}

		for _, responseParam := range endpoint.ResponseParams {
			responseParamsSection := util.RESPONSE_PARAM

			generateReadmeFileServiceImplementation.setTemplateValue("[FIELD]", responseParam.Field, &responseParamsSection)
			generateReadmeFileServiceImplementation.setTemplateValue("[TYPE]", responseParam.Type, &responseParamsSection)
			generateReadmeFileServiceImplementation.setTemplateValue("[RESPONSE_PARAM_DESCRIPTION]", responseParam.Description, &responseParamsSection)

			tempBody += responseParamsSection
		}

		body += tempBody
	}

	return header + body
}

func (generateReadmeFileServiceImplementation *generateReadmeFileServiceImplementation) setTemplateValue(key string, value string, template *string) {
	*template = strings.Replace(*template, key, value, 1)
}
