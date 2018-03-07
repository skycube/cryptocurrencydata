package cryptocurrencydata

import (
	"os"
)

var ccd *cryptocurrencydata

const (
	noDataError  = "no data or invalid currency"
	noValueError = "no value to parse"
)

func init() {
	// Default parameters
	// Default market fetch url
	useMarketURL := "https://api.coinmarketcap.com/v1/ticker/"
	// Default market url base for fetching a single currency history
	useMarketHistoryURL := "https://graphs2.coinmarketcap.com/currencies/"
	// Set value to 0 on parse error
	doSkipParseError := true
	// Parse environment variables
	if os.Getenv("CCDATA_MARKET_URL") != "" {
		useMarketURL = os.Getenv("CCDATA_MARKET_URL")
	}
	if os.Getenv("CCDATA_MARKET_HISTORY_URL") != "" {
		useMarketHistoryURL = os.Getenv("CCDATA_MARKET_HISTORY_URL")
	}
	if os.Getenv("CCDATA_SKIP_PARSE_ERROR") == "no" {
		doSkipParseError = false
	}
	// ccd create a new configuration set
	ccd = &cryptocurrencydata{
		marketURL:        useMarketURL,
		marketHistoryURL: useMarketHistoryURL,
		skipParseError:   doSkipParseError,
	}
}
