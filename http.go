package nse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/sanesource/nse-fetch/utils"
)

type NSEApiCallOptions struct {
	beforeApiFetchUrl string
}

func request(method, url string, options NSEApiCallOptions) (map[string]any, error) {
	// Step 1: Setup http client with cookie jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 10 * time.Second,
	}

	// Step 2: Visit beforeApiFetchUrl
	if len(options.beforeApiFetchUrl) > 0 {
		preApiFetchReq, _ := http.NewRequest("GET", options.beforeApiFetchUrl, nil)
		preApiFetchReq.Header.Set("User-Agent", utils.GetRandomUserAgent())
		preApiFetchReq.Header.Set("Content-Type", "text/html; charset=utf-8")
		resp, err := client.Do(preApiFetchReq)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}

	// Step 3: API Request
	req, reqErr := http.NewRequest(method, url, nil)
	if reqErr != nil {
		return nil, fmt.Errorf("failed to fetch: %s", url)
	}
	req.Header.Set("User-Agent", utils.GetRandomUserAgent())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://www.nseindia.com/")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	res, resErr := client.Do(req)
	if resErr != nil {
		return nil, fmt.Errorf("failed to fetch from nse: %s", url)
	}
	defer res.Body.Close()

	// Step 4: Read API Response
	data, readResponseErr := io.ReadAll(res.Body)
	if readResponseErr != nil {
		return nil, fmt.Errorf("failed to parse response from nse: %s", url)
	}

	var result map[string]any
	json.Unmarshal(data, &result)

	return map[string]any{"data": result}, nil
}
