package cryptostats

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const apiURL = "https://api.coinmarketcap.com/v1/ticker/"

// CoinStats represents an object for storing information about a specific coin at a set time
type CoinStats struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Symbol              string `json:"symbol"`
	Rank                string `json:"rank"`
	PriceUSD            string `json:"price_usd"`
	PrictBTC            string `json:"price_btc"`
	PercentChange24Hour string `json:"percent_change_24h"`
	PercentChange7Day   string `json:"percent_change_7d"`
	LastUpdated         string `json:"last_updated"`
}

// GetTopCoins retrieves the top n coins
func (ac *APIClient) GetTopCoins(limit int) ([]CoinStats, error) {
	if limit < 1 {
		log.Fatal("Must get at least 1 cryptocurrency")
	}
	var stats []CoinStats
	u, _ := url.Parse(apiURL)
	query := u.Query()
	query.Set("limit", strconv.Itoa(limit))
	u.RawQuery = query.Encode()
	if err := ac.Do(http.MethodGet, u.String(), nil, &stats); err != nil {
		return nil, err
	}
	return stats, nil
}

// GetCoinStats retrieves the stats for a single coin
func (ac *APIClient) GetCoinStats(coin string) ([]CoinStats, error) {
	var stats []CoinStats
	urlString := fmt.Sprintf("%s%s/", apiURL, coin)
	if err := ac.Do(http.MethodGet, urlString, nil, &stats); err != nil {
		return nil, err
	}
	return stats, nil
}
