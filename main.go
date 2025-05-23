package nse

import (
	"github.com/sanesource/nse-fetch/utils"
)

func FetchAutoComplete(symbol string) (map[string]any, error) {
	url := utils.GetNSEAutoCompletApiUrl(symbol)
	data, err := request("GET", url, NSEApiCallOptions{})
	return data, err
}

func FetchEquityHistorical(symbol, from, to string) (map[string]any, error) {
	url := utils.GetNSEHistoricalApiUrl(symbol, from, to)
	beforeApiFetchUrl := utils.GetNSEHistoricalPrefetchUrl(symbol)

	data, err := request("GET", url, NSEApiCallOptions{beforeApiFetchUrl: beforeApiFetchUrl})
	return data, err
}
