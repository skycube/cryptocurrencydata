# Cryptocurrencydata
[![Build Status](https://travis-ci.org/skycube/cryptocurrencydata.svg?branch=master)](https://travis-ci.org/skycube/cryptocurrencydata)
[![GoDoc](https://godoc.org/github.com/skycube/cryptocurrencydata?status.svg)](https://godoc.org/github.com/skycube/cryptocurrencydata)
[![Go Report Card](https://goreportcard.com/badge/github.com/skycube/cryptocurrencydata)](https://goreportcard.com/report/github.com/skycube/cryptocurrencydata)

Provides a golang library to receive current and historic data 
from coinmarketcap's API.

Please note that there limitations in how many request you can do
to their API. Please find details here: https://coinmarketcap.com/api/

**THIS IS A LIBRARY, NOT YOUR WISHING SOLUTION TO MAKE YOU RICH ALTHROUGHT IT MAY HELPS YOU TO GET THERE!**
_if you do become rich, please send me all your wins ;)_

## Why?
This library is created out of interest in playing with cryptocurrencies and
go (lang). As of writing test and playgrounds for wallets, bots and advisories
I thought this might be useful to share.
I do not consider this library as perfect, conclusive or even pretty but it works awesome.
If you have suggestions, comments or improvements or even start using it, please let me know :)

### Install
```bash
go get -u github.com/skycube/cryptocurrencydata
```

### Basic Usage

The below will show a basic usage of some of the available methods.
All examples can be found in:
Basic examples can be found here:
 
```examples/examples.go```

#### Get Value in USD for "bitcoin,ethereum,skycoin" currencies
```go
d, err := cryptocurrencydata.PriceUSDByIDs([]string{"bitcoin", "ethereum", "skycoin"})
if err != nil {
    fmt.Printf("%v", err)
}
fmt.Println("# PriceUSDByIDs")
for _, v := range d {
    fmt.Printf("%-10s \t $%.2f\n", v.ID, v.Value)
}
```
Output:
```bash
# GetPriceUSDByCurrencyIDs
bitcoin    	 $10638.00
ethereum   	 $787.79
skycoin    	 $14.02
```

#### Multiple currencies values change last 1 hour
```go
d, err := cryptocurrencydata.GetPercentChange1HByCurrencyIDs([]string{"bitcoin", "ethereum", "skycoin"})
if err != nil {
    fmt.Printf("%v", err)
}
fmt.Println("# GetPercentChange1HByCurrencyIDs")
for _, v := range d {
    fmt.Printf("%-10s \t %.2f\n", v.ID, v.Value)
}
```
Output:
```bash
# GetPercentChange1HByCurrencyIDs
bitcoin    	 0.55
ethereum   	 0.31
skycoin    	 2.50
```

#### Get full currencies data details by currency ids
```go
d, err := cryptocurrencydata.GetCurrencyDataByIDs([]string{"bitcoin", "ethereum", "skycoin"})
if err != nil {
    fmt.Printf("%v", err)
}
fmt.Println("# GetCurrencyDataByIDs")
for _, v := range d {
    fmt.Printf("%v\n", v)
}
```
Output:
```bash
# GetCurrencyDataByCurrencyIDs
&{bitcoin Bitcoin BTC 1 10638 1 6.91178e+09 1.79837645256e+11 1.6905212e+07 1.6905212e+07 2.1e+07 0.55 -3.61 -0.32 2018-03-07 11:39:26 +0000 UTC}
&{ethereum Ethereum ETH 2 787.794 0.074382 1.86455e+09 7.7233590834e+10 9.80378e+07 9.80378e+07 0 0.31 -5.76 -9.69 2018-03-07 11:39:12 +0000 UTC}
&{skycoin Skycoin SKY 120 14.0169 0.00132345 552055 1.07331524e+08 7.657294e+06 2.5e+07 1e+08 2.5 -5.5 -16.3 2018-03-07 11:39:10 +0000 UTC}
```

#### Get single currency history
```go
d, err := cryptocurrencydata.GetCurrencyHistoryByCurrencyID("skycoin")
if err != nil {
    fmt.Printf("%v", err)
}
fmt.Println("# GetCurrencyHistoryByCurrencyID")
for _, v := range d {
    fmt.Printf("timestamp %v timeUTC: %v MarketSupply: %d PriceUSD: $%.2f PriceBTC: $%.2f VolUSD: $%d\n", v.Timestamp, v.TimeUTC, v.MarketSupply, v.PriceUSD, v.PriceBTC, v.VolUSD)
}
```
Output:
```bash
# GetCurrencyHistoryByCurrencyID
timestamp 1519029250000 timeUTC: 2018-02-19 08:34:10 +0000 UTC MarketSupply: 147280561 PriceUSD: $19.76 PriceBTC: $0.00 VolUSD: $670990
timestamp 1496298589000 timeUTC: 2017-06-01 06:29:49 +0000 UTC MarketSupply: 8866853 PriceUSD: $1.63 PriceBTC: $0.00 VolUSD: $73770
timestamp 1507642758000 timeUTC: 2017-10-10 13:39:18 +0000 UTC MarketSupply: 17325056 PriceUSD: $2.92 PriceBTC: $0.00 VolUSD: $9757
timestamp 1509521956000 timeUTC: 2017-11-01 07:39:16 +0000 UTC MarketSupply: 24330265 PriceUSD: $4.10 PriceBTC: $0.00 VolUSD: $33590
timestamp 1509845958000 timeUTC: 2017-11-05 01:39:18 +0000 UTC MarketSupply: 24284342 PriceUSD: $4.09 PriceBTC: $0.00 VolUSD: $13874
timestamp 1513735163000 timeUTC: 2017-12-20 01:59:23 +0000 UTC MarketSupply: 107126489 PriceUSD: $16.92 PriceBTC: $0.00 VolUSD: $320845
...
```

## Todo
* Optional parameters
* Code documentation (yes I know...)
* SpelLChaeck
* Write tests one close to publish (just hate them atm)
* Suggestions?

### Things may going to happen
* Do things better, prettier, writing GO is always learning and improving
* Additional data feeds
* Configuration/ pass in URLs (i.e. override incoming with proxy)

### Things not going to happen
* Trotteling number of requests
(as remote API limits the number of request;  10 per minute)
* Single Currency lookups, can't find a use case for this
* Store (local) to prevent duplicate lookups
* Support more than one thread ad a time
* Proxy responses internaly
* Logging
* Abstractions (this is not chinese JAVA University project)

### Contribute
* Please open an issue first
* Create a pull request
* Document your change
* go vet, go imports, go fmt and all other usual suspects first
* Don't miss changing existing tests
* Be patient, it will maybe take 1h until I have time to check your change

#### Create an Issue First

Your PR is more likely to be accepted if you open an issue first.

#### Submission Checklist

- [ ] `gofmt -s -l .` produces no output
- [ ] `go vet .` produces no output
- [ ] `golint ./...` produces no output
- [ ] `misspell .` produces no output
- [ ] `ineffassign .` produces no output
- [ ] `go test -race` passes

## Thanks
Thanks to everyone reading this, using this and spcialy to those who are 
not just negative.
Be kind, do kind, help, not be a doushbag!