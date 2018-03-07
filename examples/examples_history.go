package main

import (
	"fmt"

	ccdata "github.com/skycube/cryptocurrencydata"
)

func main() {

	// Get single currency history
	currencyHistory, err := ccdata.GetCurrencyHistoryByCurrencyID("skycoin")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetCurrencyHistoryByCurrencyID")
	for _, v := range currencyHistory {
		fmt.Printf("timestamp %v timeUTC: %v MarketSupply: %d PriceUSD: $%.2f PriceBTC: $%.2f VolUSD: $%d\n", v.Timestamp, v.TimeUTC, v.MarketSupply, v.PriceUSD, v.PriceBTC, v.VolUSD)
	}

}
