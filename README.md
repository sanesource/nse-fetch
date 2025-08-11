# nse-fetch

Fetch indian stock market data available on nse using this library

## Installation

```bash
go get -u github.com/sanesource/nse-fetch
```

## Methods

- #### FetchAutoComplete

Takes a **symbol** (_string_) keyword, and returns **instruments** matching that keyword

<details>
<summary>
Example
</summary>

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

</details>

- #### FetchEquityHistorical

Takes **symbol** (_string_), **from** (_string_) and **to** (_string_) and returns historical data for given symbol

<details>
<summary>
Example
</summary>

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

</details>

- #### FetchNifty50Performance

Returns Nifty 50 index and its 50 constituents latest performance data

<details>
<summary>
Example
</summary>

```go
// Import package
import (
	NSE "github.com/sanesource/nse-fetch"
)

// Usage
func main() {
    response, err := NSE.FetchNifty50Performance()
    if err != nil {
        // handle error here
    }

    // response -> { "data": [...] }
}
```

</details>

## Status

Currently WIP
