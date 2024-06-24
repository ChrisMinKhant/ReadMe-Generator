package httpsfilter

import (
	"net/http"

	"github.com/ChrisMinKhant/megoyougo_framework/filter"
	"github.com/sirupsen/logrus"
)

type httpsFilter struct {
}

func New() *httpsFilter {
	return &httpsFilter{}
}

func (httpsFilter *httpsFilter) Do(response http.ResponseWriter, request *http.Request) {
	if request.TLS == nil {

		logrus.Error("Request is not from https.")
		filter.ErrorSigal <- " HTTPS FILTER FAILED "
		return

	}
	filter.ErrorSigal <- ""
}
