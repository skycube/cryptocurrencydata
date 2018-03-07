package main

import (
	"fmt"

	"github.com/skycube/cryptocurrencydata"
)

func main() {

	// Multiple currencies price in USD
	cValuesUSD, err := cryptocurrencydata.GetPriceUSDByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetPriceUSDByCurrencyIDs")
	for _, v := range cValuesUSD {
		fmt.Printf("%-10s \t $%.2f\n", v.ID, v.Value)
	}

	// Multiple currencies values change last 1 hour
	cValuesChange1h, err := cryptocurrencydata.GetPercentChange1HByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetPercentChange1HByCurrencyIDs")
	for _, v := range cValuesChange1h {
		fmt.Printf("%-10s \t %.2f\n", v.ID, v.Value)
	}

	// Multiple currencies full data
	cMCData, err := cryptocurrencydata.GetCurrencyDataByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetCurrencyDataByCurrencyIDs")
	for _, v := range cMCData {
		fmt.Printf("%v\n", v)
	}

	// Get single currency history
	currencyHistory, err := cryptocurrencydata.GetCurrencyHistoryByCurrencyID("skycoin")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetCurrencyHistoryByCurrencyID")
	for _, v := range currencyHistory {
		fmt.Printf("timestamp %v timeUTC: %v MarketSupply: %d PriceUSD: $%.2f PriceBTC: $%.2f VolUSD: $%d\n", v.Timestamp, v.TimeUTC, v.MarketSupply, v.PriceUSD, v.PriceBTC, v.VolUSD)
	}

}
