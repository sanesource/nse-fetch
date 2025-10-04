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

func FetchNifty50Performance() (map[string]any, error) {
	url := utils.GetNifty50PerformaceUrl()
	beforeApiFetchUrl := utils.GetNifty50PerformacePrefetchUrl()
	data, err := request("GET", url, NSEApiCallOptions{beforeApiFetchUrl: beforeApiFetchUrl})
	return data, err
}

func FetchTopGainers() (map[string]any, error) {
	url := utils.GetTopGainersUrl()
	beforeApiFetchUrl := utils.GetTopGainersLosersPrefetchUrl()
	data, err := request("GET", url, NSEApiCallOptions{beforeApiFetchUrl: beforeApiFetchUrl})
	return data, err
}

func FetchTopLosers() (map[string]any, error) {
	url := utils.GetTopLosersUrl()
	beforeApiFetchUrl := utils.GetTopGainersLosersPrefetchUrl()
	data, err := request("GET", url, NSEApiCallOptions{beforeApiFetchUrl: beforeApiFetchUrl})
	return data, err
}
