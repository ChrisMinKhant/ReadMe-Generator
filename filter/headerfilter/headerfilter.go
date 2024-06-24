package headerfilter

import (
	"net/http"
	"strings"

	"github.com/ChrisMinKhant/megoyougo_framework/filter"
	"github.com/sirupsen/logrus"
)

type headerFilter struct {
}

func New() *headerFilter {
	return &headerFilter{}
}

func (headerFilter *headerFilter) Do(response http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {

		logrus.Error("Request header doesn't contain [ Content-Type : application/json ]")
		filter.ErrorSigal <- " HEADER FILTER FAILED "
		return

	}
	filter.ErrorSigal <- ""
}
