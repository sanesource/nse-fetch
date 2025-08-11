package utils

import "fmt"

func GetNSEAutoCompletApiUrl(symbol string) string {
	return fmt.Sprintf("https://www.nseindia.com/api/search/autocomplete?q=%s", symbol)
}

func GetNSEHistoricalPrefetchUrl(symbol string) string {
	return fmt.Sprintf("https://www.nseindia.com/get-quotes/equity?symbol=%s", symbol)
}

func GetNSEHistoricalApiUrl(symbol, from, to string) string {
	return fmt.Sprintf("https://www.nseindia.com/api/historical/cm/equity?symbol=%s&series=[\"EQ\"]&from=%s&to=%s", symbol, from, to)
}

func GetNifty50PerformaceUrl() string {
	return "https://www.nseindia.com/api/equity-stockIndices?index=NIFTY%2050"
}

func GetNifty50PerformacePrefetchUrl() string {
	return "https://www.nseindia.com/market-data/top-gainers-losers"
}
