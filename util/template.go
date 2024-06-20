package util

var HEADER string = "# [ServiceName] Service [ServiceCategory]\n" +
	"### Service Description\n" +
	"[ServiceDescription]\n"

var Endpoint string = "#### Endpoint [EndpointNumber]\n" +
	"- [Path]\n" +
	"- #### Request Body\n" +
	"````json\n" +
	"    [ExampleRequest]\n" +
	"````\n" +
	"- #### Respone Body\n" +
	"````json \n" +
	"    [ExampleResponse]\n" +
	"````\n"

var REQUEST_PARAM = "#### Request Parameters\n" +
	"##### [FIELD] [TYPE]\n" +
	"##### Validations\n" +
	"| Name | Description |\n" +
	"| :---: | :---: |\n" +
	"[VALIDATIONS]" +
	"##### Description\n" +
	"[REQUEST_PARAM_DESCRIPTION]\n"

var RESPONSE_PARAM = "#### Response Parameters\n" +
	"##### [FIELD] [TYPE]\n" +
	"##### Description\n" +
	"[RESPONSE_PARAM_DESCRIPTION]"
