# nse-fetch

Fetch indian stock market data available on nse using this library

## Installation

```bash
go get -u github.com/sanesource/nse-fetch
```

## Methods

- ### FetchAutoComplete

Takes a **symbol** keyword, and returns **instruments** matching that keyword

```go
// Import package
import (
	NSE "github.com/sanesource/nse-fetch"
)

// Usage
func main() {
    response, err := NSE.FetchAutoComplete(symbol)
    if err != nil {
        // handle error here
    }

    // response -> { "data": ... }
}
```

- #### FetchEquityHistorical

Takes **symbol** (_string_), **from** (_string_) and **to** (_string_) and returns historical data for given symbol

```go
// Import package
import (
	NSE "github.com/sanesource/nse-fetch"
)

// Usage
func main() {
    response, err := NSE.FetchEquityHistorical(symbol, "23-05-2025", "23-05-2025")
    if err != nil {
        // handle error here
    }

    // response -> { "data": [...] }
}
```

## Status

Currently WIP
