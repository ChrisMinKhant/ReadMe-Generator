package filterchain

import (
	"net/http"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"github.com/ChrisMinKhant/megoyougo_framework/filter/headerfilter"
	"github.com/ChrisMinKhant/megoyougo_framework/filter/whitelistfilter"
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
	filterChain.filterList.Add(headerfilter.New())
	filterChain.filterList.Add(whitelistfilter.New())
}

func (filterChain *FilterChain) Invoke(response http.ResponseWriter, request *http.Request) bool {
	return filterChain.filterList.Invoke(response, request)
}
