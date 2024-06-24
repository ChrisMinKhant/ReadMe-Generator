package tokenfilter

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/filter"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type tokenFilter struct {
	token     map[string]string
	exception exception.Exception
}

var tokenFilterInstance *tokenFilter

func New() *tokenFilter {
	if tokenFilterInstance != nil {
		return tokenFilterInstance
	}

	tokenFilterInstance := &tokenFilter{
		token:     make(map[string]string),
		exception: exception.GetGeneralExceptionInstance(),
	}

	return tokenFilterInstance
}

func (tokenFilter *tokenFilter) Do(response http.ResponseWriter, request *http.Request) {
	defer tokenFilter.exception.RecoverPanic()

	requestIp := tokenFilter.fetchIp(request.RemoteAddr)

	if tokenFilter.token[requestIp] != "" {
		var decodedToken []byte

		logrus.Infof("Fetched app-token ::: %v\n", request.Header.Get("app-token"))

		_, error := base64.StdEncoding.Decode(decodedToken, []byte(request.Header.Get("app-token")))

		logrus.Infof("Fetched decoded token ::: %v\n", string(decodedToken))

		if error != nil {
			logrus.Panicf("Error occurred at decoding token ::: %v\n", error.Error())
		}

		if tokenFilter.token[requestIp] != string(decodedToken) {
			filter.ErrorSigal <- " TOKEN FILTER FAILED "
		}
		filter.ErrorSigal <- ""
	}

	tokenFilter.token[requestIp] = requestIp + "-" + uuid.NewString()

	filter.ErrorSigal <- " TOKEN FILTER FAILED - Please use this token on next request => " + base64.StdEncoding.EncodeToString([]byte(tokenFilter.token[requestIp]))
}

func (tokenFilter *tokenFilter) fetchIp(rawRequestOrigin string) string {
	firstProcessing := strings.Split(rawRequestOrigin, ":")
	secondProcessing := strings.Split(firstProcessing[2], "]")

	return secondProcessing[0]
}
