package cryptocurrencydata

import "time"

type cryptocurrencydata struct {
	marketURL        string
	marketHistoryURL string
	skipParseError   bool
}

type coinmarketcapcurrency struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUSD         string `json:"price_usd"`
	PriceBTC         string `json:"price_btc"`
	VolumeUSD24H     string `json:"24h_volume_usd"`
	MarketCapUSD     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

// Cyptocurrency is the type of a current currency
type Cyptocurrency struct {
	ID               string
	Name             string
	Symbol           string
	Rank             int64
	PriceUSD         float64
	PriceBTC         float64
	VolumeUSD24H     float64
	MarketCapUSD     float64
	AvailableSupply  float64
	TotalSupply      float64
	MaxSupply        float64
	PercentChange1H  float64
	PercentChange24H float64
	PercentChange7D  float64
	LastUpdated      time.Time
}

// ValueResponse is the single value response type
type ValueResponse struct {
	ID    string
	Value float64
}

type coinmarketcaphistory struct {
	MarketSupply [][]interface{} `json:"market_cap_by_available_supply"`
	PriceBTC     [][]interface{} `json:"price_btc"`
	PriceUSD     [][]interface{} `json:"price_usd"`
	VolUSD       [][]interface{} `json:"volume_usd"`
}

// CryptocurrencyHistory is the type of a historic currency
type CryptocurrencyHistory struct {
	Timestamp    int64
	TimeUTC      time.Time
	MarketSupply int64
	PriceBTC     float64
	PriceUSD     float64
	VolUSD       int64
}
