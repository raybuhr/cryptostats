package main

import (
	"log"

	"github.com/raybuhr/cryptostats"
)

func main() {
	coinmarketcap := cryptostats.New("none")
	// BitShares aka bts
	bts, err := coinmarketcap.GetCoinStats("bitshares")
	if err != nil {
		log.Println(err)
	}
	log.Printf("%v\n", bts)

	topFive, err := coinmarketcap.GetTopCoins(5)
	if err != nil {
		log.Println(err)
	}
	for _, coin := range topFive {
		log.Println(coin)
	}
}
