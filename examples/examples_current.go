package main

import (
	"fmt"

	ccdata "github.com/skycube/cryptocurrencydata"
)

func main() {

	// Multiple currencies price in USD
	cValuesUSD, err := ccdata.GetPriceUSDByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetPriceUSDByCurrencyIDs")
	for _, v := range cValuesUSD {
		fmt.Printf("%-10s \t $%.2f\n", v.ID, v.Value)
	}

	// Multiple currencies values change last 1 hour
	cValuesChange1h, err := ccdata.GetPercentChange1HByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetPercentChange1HByCurrencyIDs")
	for _, v := range cValuesChange1h {
		fmt.Printf("%-10s \t %.2f\n", v.ID, v.Value)
	}

	// Multiple currencies full data
	cMCData, err := ccdata.GetCurrencyDataByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println("# GetCurrencyDataByCurrencyIDs")
	for _, v := range cMCData {
		fmt.Printf("%v\n", v)
	}

}
