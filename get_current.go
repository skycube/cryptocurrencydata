package cryptocurrencydata

import "errors"

// getSingleValueByCurrencyIDs receives a single value type by ids and name/type
func getSingleValueByCurrencyIDs(ids []string, valueType string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	data, err := fetchMarketCurrencyByIDs(ids)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New(noDataError)
	}
	for _, v := range data {
		value, err := getValue(valueType, v)
		if err != nil {
			return nil, err
		}
		rsp = append(rsp, value)
	}
	return rsp, nil
}

// GetPriceUSDByCurrencyIDs get current price in USD for given IDs
func GetPriceUSDByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "PriceUSD")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetPriceBTCByCurrencyIDs get current price in BTC for given IDs
func GetPriceBTCByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "PriceBTC")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetVolumeUSD24HByCurrencyIDs get current volume over last 24 hours in USD for given IDs
func GetVolumeUSD24HByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "VolumeUSD24H")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetMarketCapUSDByCurrencyIDs get current market cap in USD for given IDs
func GetMarketCapUSDByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "MarketCapUSD")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetAvailableSupplyByCurrencyIDs get current available supply for given IDs
func GetAvailableSupplyByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "AvailableSupply")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetTotalSupplyByCurrencyIDs get current toal supply for given IDs
func GetTotalSupplyByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "TotalSupply")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetMaxSupplyByCurrencyIDs get current max supply for given IDs
func GetMaxSupplyByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "MaxSupply")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetPercentChange1HByCurrencyIDs get change in percentage over last 1 hour for given IDs
func GetPercentChange1HByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "PercentChange1H")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetPercentChange24HByCurrencyIDs get change in percentage over last 24 hours for given IDs
func GetPercentChange24HByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "PercentChange24H")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetPercentChange7DHByCurrencyIDs get change in percentage over last 7 days for given IDs
func GetPercentChange7DHByCurrencyIDs(ids []string) ([]*ValueResponse, error) {
	var rsp []*ValueResponse
	rsp, err := getSingleValueByCurrencyIDs(ids, "PercentChange7D")
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// GetCurrencyDataByCurrencyIDs get all current details about a given crypto currency
func GetCurrencyDataByCurrencyIDs(ids []string) ([]*Cyptocurrency, error) {
	var rsp []*Cyptocurrency
	data, err := fetchMarketCurrencyByIDs(ids)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New(noDataError)
	}
	rsp = data
	return rsp, nil
}
