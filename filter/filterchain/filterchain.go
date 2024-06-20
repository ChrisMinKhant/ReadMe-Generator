package filterchain

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/filter"
	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"github.com/sirupsen/logrus"
)

type FilterChain struct {
	filterList filterList
	exception  exception.Exception
}

func New() *FilterChain {
	return &FilterChain{
		filterList: *NewFilterList(),
		exception:  exception.GetGeneralExceptionInstance(),
	}
}

/*
 * The filter must be listed here.
 * All the filter that are existing here
 * will be invoked on each http request.
 */

func (filterChain *FilterChain) Set() {

	/*
	 * The filters will be invoked in top to bottom
	 * order.
	 */

	// filterChain.filterList.Add(httpsfilter.New())
	// filterChain.filterList.Add(headerfilter.New())
	// filterChain.filterList.Add(authfilter.New())
}

func (filterChain *FilterChain) Invoke(response http.ResponseWriter, request *http.Request) bool {
	defer filterChain.exception.RecoverPanic()

	filterChain.filterList.Invoke(response, request)

	if fetchedSignal := <-filter.ErrorSigal; fetchedSignal != "" {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(util.NewErrorResponse().SetStatus("Filteration Failed").SetMessage(fetchedSignal).SetPath(request.RequestURI).SetTimestamp(time.Now().String()))

		logrus.Panicf("Filteration failed with error ::: [ %v ]\n", fetchedSignal)

		return false
	}

	return true
}
