# Cryptocurrencydata
[![Build Status](https://travis-ci.org/skycube/cryptocurrencydata.svg?branch=develop)](https://travis-ci.org/skycube/cryptocurrencydata)
[![GoDoc](https://godoc.org/github.com/skycube/cryptocurrencydata?status.svg)](https://godoc.org/github.com/skycube/cryptocurrencydata)
[![Go Report Card](https://goreportcard.com/badge/github.com/skycube/cryptocurrencydata)](https://goreportcard.com/report/github.com/skycube/cryptocurrencydata)

Provides a golang library to receive current and historic data 
from coinmarketcap's API.

Please note that there limitations in how many request you can do
to their API. Please find details here https://coinmarketcap.com/api/

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

Get Value in USD for "bitcoin,ethereum,skycoin" currencies
```go
// Multiple currencies price in USD
d, err := ccd.PriceUSDByIDs([]string{"bitcoin", "ethereum", "skycoin"})
if err != nil {
    fmt.Printf("%v", err)
}
fmt.Println("# PriceUSDByIDs")
for _, v := range d {
    fmt.Printf("%-10s \t $%.2f\n", v.ID, v.Value)
}
```

Get full currencies data details by currency ids
```go
// Multiple currencies full data
d, err := ccd.GetCurrencyDataByIDs([]string{"bitcoin", "ethereum", "skycoin"})
if err != nil {
    fmt.Printf("%v", err)
}
fmt.Println("# GetCurrencyDataByIDs")
for _, v := range d {
    fmt.Printf("%v\n", v)
}
```

### Examples

Can be found here:
 
```examples/*```

## Todo
* Optional parameters
* Code documentation (yes I know...)
* SpelLChaeck
* Write tests one close to publish (just hate them atm)
* ...something else?...

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