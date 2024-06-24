package whitelistfilter

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/filter"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type whiteListFilter struct {
	tokens    map[string]string
	exception exception.Exception
}

var whiteListFilterInstance *whiteListFilter

func New() *whiteListFilter {
	if whiteListFilterInstance != nil {
		return whiteListFilterInstance
	}

	whiteListFilterInstance := &whiteListFilter{
		tokens:    make(map[string]string),
		exception: exception.GetGeneralExceptionInstance(),
	}

	return whiteListFilterInstance
}

func (whiteListFilter *whiteListFilter) Do(response http.ResponseWriter, request *http.Request) {
	defer whiteListFilter.exception.RecoverPanic()

	requestIp := whiteListFilter.fetchIp(request.RemoteAddr)

	// Check if the ip is white listed
	for _, whiteListedIp := range WhiteList {

		if requestIp == whiteListedIp {

			// Check if there is already a token for this ip
			if whiteListFilter.tokens[requestIp] != "" {

				logrus.Infof("Token was already generated for ip [ %v ].\n", requestIp)

				decodedTokenByte, error := base64.StdEncoding.DecodeString(request.Header.Get("app-token"))

				logrus.Infof("Fetched decoded token ::: %v\n", string(decodedTokenByte))

				if error != nil {
					logrus.Panicf("Error occurred at decoding token ::: %v\n", error.Error())
				}

				if whiteListFilter.tokens[requestIp] != string(decodedTokenByte) {
					filter.ErrorSigal <- " TOKEN FILTER FAILED "
					return
				}

				filter.ErrorSigal <- ""
				return
			}

			logrus.Infof("There is no generated token for ip [ %v ].\n", requestIp)
			logrus.Infof("Generating token for ip [ %v ].\n", requestIp)
			// If the token doesn't exist, generate one for this ip
			whiteListFilter.tokens[requestIp] = requestIp + "-" + uuid.NewString()

			filter.ErrorSigal <- " TOKEN FILTER FAILED - Please use this token on next request => " + base64.StdEncoding.EncodeToString([]byte(whiteListFilter.tokens[requestIp]))
			return
		}
	}

	logrus.Errorf("The requested ip [ %v ] is not white listed.\n", requestIp)

	filter.ErrorSigal <- " TOKEN FILTER FAILED - This request was comming from unknown member or outside of white listed member."
}

func (whiteListFilter *whiteListFilter) fetchIp(rawRequestOrigin string) string {
	firstProcessing := strings.Split(rawRequestOrigin, ":")
	secondProcessing := strings.Split(firstProcessing[2], "]")

	return secondProcessing[0]
}
