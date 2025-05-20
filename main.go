package nse

import "fmt"

func FetchAutoComplete(symbol string) (map[string]any, error) {
	url := fmt.Sprintf("https://www.nseindia.com/api/search/autocomplete?q=%s", symbol)
	data, err := makeNSEApiCall("GET", url)
	return data, err
}
