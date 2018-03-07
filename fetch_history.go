package cryptocurrencydata

import (
	"encoding/json"
	"fmt"

	"net/http"

	"time"
)

// fetchCurrencyHistoryByCurrencyID
func fetchCurrencyHistoryByCurrencyID(id string) (map[int64]*CryptocurrencyHistory, error) {
	historicData := map[int64]*CryptocurrencyHistory{}

	fetchURL := fmt.Sprintf("%s%s", ccd.marketHistoryURL, id)
	response, err := http.Get(fetchURL)
	data := coinmarketcaphistory{}
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	// Init and add MarketSupply
	for _, v := range data.MarketSupply {
		theTimestamp := int64(v[0].(float64))
		rTimestamp := theTimestamp / 1000
		theValue := int64(v[1].(float64))
		entry := &CryptocurrencyHistory{
			Timestamp:    theTimestamp,
			TimeUTC:      time.Unix(rTimestamp, 0).UTC(),
			MarketSupply: theValue,
		}
		historicData[theTimestamp] = entry
	}

	// Add PriceBtc
	for _, v := range data.PriceBTC {
		theTimestamp := int64(v[0].(float64))
		theValue := v[1].(float64)
		historicData[theTimestamp].PriceBTC = theValue
	}

	// Add PriceUsd
	for _, v := range data.PriceUSD {
		theTimestamp := int64(v[0].(float64))
		theValue := v[1].(float64)
		historicData[theTimestamp].PriceUSD = theValue
	}

	// Add VolUsd
	for _, v := range data.VolUSD {
		theTimestamp := int64(v[0].(float64))
		theValue := v[1].(float64)
		historicData[theTimestamp].VolUSD = int64(theValue)
	}

	//spew.Dump(historicData)

	return historicData, nil
}
