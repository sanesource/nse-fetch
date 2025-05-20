package nse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sanesource/nse-fetch/utils"
)

func makeNSEApiCall(method, url string) (map[string]any, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, fmt.Errorf("failed to fetch: %s", url)
	}

	userAgent := utils.GetRandomUserAgent()
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.nseindia.com/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, fmt.Errorf("failed to fetch from nse: %s", url)
	}
	defer res.Body.Close()

	data, readResponseErr := io.ReadAll(res.Body)
	if readResponseErr != nil {
		return nil, fmt.Errorf("failed to parse response from nse: %s", url)
	}

	var result map[string]any
	json.Unmarshal(data, &result)

	return map[string]any{"data": result}, nil
}
