package cryptocurrencydata

// GetCurrencyHistoryByCurrencyID get historic data for a currency by ID
func GetCurrencyHistoryByCurrencyID(id string) (map[int64]*CryptocurrencyHistory, error) {
	data, err := fetchCurrencyHistoryByCurrencyID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
