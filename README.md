<div align="center">

# 📈 nse-fetch

**A lightweight Go library for fetching Indian stock market data from NSE**

[![Go Version](https://img.shields.io/badge/Go-1.24.2+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/sanesource/nse-fetch)](https://goreportcard.com/report/github.com/sanesource/nse-fetch)

[Features](#-features) • [Installation](#-installation) • [API Reference](#-api-reference) • [Contributing](#-contributing)

</div>

---

## 🎯 Overview

`nse-fetch` is a simple and efficient Go library that provides programmatic access to the National Stock Exchange of India (NSE) market data. Fetch real-time stock information, historical data, and market indices with just a few lines of code.

## ✨ Features

- 🔍 **Auto-complete Search** - Find stock symbols and instruments quickly
- 📊 **Historical Data** - Fetch historical equity data for any date range
- 📈 **Nifty 50 Performance** - Get real-time data for Nifty 50 index and constituents
- 🚀 **Zero Dependencies** - Uses only Go standard library
- ⚡ **Fast & Lightweight** - Optimized for performance
- 🔒 **Cookie Handling** - Automatically manages NSE session cookies

## 📦 Installation

```bash
go get -u github.com/sanesource/nse-fetch
```

**Requirements:**

- Go 1.24.2 or higher

## 📚 API Reference

### FetchAutoComplete

Search for stock symbols and get instrument suggestions.

**Signature:**

```go
func FetchAutoComplete(symbol string) (map[string]any, error)
```

**Parameters:**

- `symbol` (string) - The search keyword or partial symbol name

**Returns:**

- `map[string]any` - Response containing matching instruments
- `error` - Error if the request fails

**Example:**

```go
package main

import (
    "fmt"
    "log"

    nse "github.com/sanesource/nse-fetch"
)

func main() {
    response, err := nse.FetchAutoComplete("TATA")
    if err != nil {
        log.Fatal(err)
    }

    // Response structure: { "data": { "symbols": [...] } }
    fmt.Printf("Found instruments: %v\n", response)
}
```

---

### FetchEquityHistorical

Retrieve historical stock data for a specific symbol and date range.

**Signature:**

```go
func FetchEquityHistorical(symbol, from, to string) (map[string]any, error)
```

**Parameters:**

- `symbol` (string) - The stock symbol (e.g., "RELIANCE", "TCS")
- `from` (string) - Start date in format `DD-MM-YYYY`
- `to` (string) - End date in format `DD-MM-YYYY`

**Returns:**

- `map[string]any` - Response containing historical data
- `error` - Error if the request fails

**Example:**

```go
package main

import (
    "fmt"
    "log"

    nse "github.com/sanesource/nse-fetch"
)

func main() {
    // Fetch historical data for RELIANCE from Jan 1 to Jan 31, 2025
    response, err := nse.FetchEquityHistorical("RELIANCE", "01-01-2025", "31-01-2025")
    if err != nil {
        log.Fatal(err)
    }

    // Response structure: { "data": [...] }
    fmt.Printf("Historical data: %v\n", response)
}
```

---

### FetchNifty50Performance

Get the latest performance data for Nifty 50 index and all its constituent stocks.

**Signature:**

```go
func FetchNifty50Performance() (map[string]any, error)
```

**Parameters:**

- None

**Returns:**

- `map[string]any` - Response containing Nifty 50 performance data
- `error` - Error if the request fails

**Example:**

```go
package main

import (
    "fmt"
    "log"

    nse "github.com/sanesource/nse-fetch"
)

func main() {
    response, err := nse.FetchNifty50Performance()
    if err != nil {
        log.Fatal(err)
    }

    // Response structure: { "data": [...] }
    fmt.Printf("Nifty 50 data: %v\n", response)
}
```

## 🛠️ Development

### Running Tests

```bash
go test -v ./...
```

### Building from Source

```bash
git clone https://github.com/sanesource/nse-fetch.git
cd nse-fetch
go build
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## ⚠️ Disclaimer

This library is for educational and research purposes only. Please ensure you comply with NSE's terms of service and usage policies when using this library. The authors are not responsible for any misuse of this library.

## 🙏 Acknowledgments

- Data provided by [National Stock Exchange of India](https://www.nseindia.com/)
- Inspired by the need for easy programmatic access to Indian stock market data

## 📧 Support

- **Issues**: [GitHub Issues](https://github.com/sanesource/nse-fetch/issues)
- **Discussions**: [GitHub Discussions](https://github.com/sanesource/nse-fetch/discussions)

---

<div align="center">

Made with ❤️ by the sanesource team

**[⬆ back to top](#-nse-fetch)**

</div>
