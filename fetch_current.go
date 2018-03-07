package cryptocurrencydata

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// fetchMarketCurrencyByIDs fetches given currencies
func fetchMarketCurrencyByIDs(ids []string) ([]*Cyptocurrency, error) {
	var rsp []*Cyptocurrency
	var fetchURL string
	if len(ids) == 1 {
		fetchURL = fmt.Sprintf("%s%s", ccd.marketURL, strings.Join(ids, ","))
	} else {
		fetchURL = fmt.Sprintf("%s?limit=0", ccd.marketURL)
	}
	res, err := http.Get(fetchURL)
	if err != nil {
		return rsp, err
	}
	defer res.Body.Close()
	var data []coinmarketcapcurrency
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return rsp, err
	}
	if data != nil {
		for _, v := range data {
			if len(ids) == 0 {
				parsed, err := parseMarketData(v)
				if err != nil {
					return rsp, err
				}
				rsp = append(rsp, parsed)
			} else {
				for _, i := range ids {
					if i == v.ID {
						parsed, err := parseMarketData(v)
						if err != nil {
							return rsp, err
						}
						rsp = append(rsp, parsed)
					}
				}
			}

		}
	}
	return rsp, nil
}

// parseMarketData parses a single fetched currency
func parseMarketData(cmc coinmarketcapcurrency) (*Cyptocurrency, error) {
	rsp := &Cyptocurrency{}
	var err error
	if cmc.Name != "" {
		rsp.ID = cmc.ID
		rsp.Name = cmc.Name
		rsp.Symbol = cmc.Symbol
		rsp.Rank, err = strconv.ParseInt(cmc.Rank, 10, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.PriceUSD, err = strconv.ParseFloat(cmc.PriceUSD, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.PriceBTC, err = strconv.ParseFloat(cmc.PriceBTC, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.VolumeUSD24H, err = strconv.ParseFloat(cmc.VolumeUSD24H, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.MarketCapUSD, err = strconv.ParseFloat(cmc.MarketCapUSD, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.AvailableSupply, err = strconv.ParseFloat(cmc.AvailableSupply, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.TotalSupply, err = strconv.ParseFloat(cmc.TotalSupply, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.MaxSupply, err = strconv.ParseFloat(cmc.MaxSupply, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.PercentChange1H, err = strconv.ParseFloat(cmc.PercentChange1H, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.PercentChange24H, err = strconv.ParseFloat(cmc.PercentChange24H, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.PercentChange7D, err = strconv.ParseFloat(cmc.PercentChange7D, 64)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		lastUpdatedI, err := strconv.Atoi(cmc.LastUpdated)
		if err != nil && !ccd.skipParseError {
			return nil, err
		}
		rsp.LastUpdated = time.Unix(int64(lastUpdatedI), 0).UTC()
	}
	return rsp, nil
}

// getValue get a certain value for a given Currency
func getValue(valueName string, d *Cyptocurrency) (*ValueResponse, error) {
	if valueName == "" || d == nil {
		return nil, errors.New(noValueError)
	}
	rsp := &ValueResponse{
		ID: d.ID,
	}
	switch valueName {
	case "PriceUSD":
		rsp.Value = d.PriceUSD
		break
	case "PriceBTC":
		rsp.Value = d.PriceBTC
		break
	case "VolumeUSD24H":
		rsp.Value = d.VolumeUSD24H
		break
	case "MarketCapUSD":
		rsp.Value = d.MarketCapUSD
		break
	case "AvailableSupply":
		rsp.Value = d.AvailableSupply
		break
	case "TotalSupply":
		rsp.Value = d.TotalSupply
		break
	case "MaxSupply":
		rsp.Value = d.MaxSupply
		break
	case "PercentChange1H":
		rsp.Value = d.PercentChange1H
		break
	case "PercentChange24H":
		rsp.Value = d.PercentChange24H
		break
	case "PercentChange7D":
		rsp.Value = d.PercentChange7D
		break
	default:
		return nil, nil
	}
	return rsp, nil
}
