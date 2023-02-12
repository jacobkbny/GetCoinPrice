package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Ticker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	WantedCoin := "XRP"
	price, err := GetCoinPrice(WantedCoin)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The current price of %s on Binance is $%s\n", WantedCoin, price)
}

func GetCoinPrice(coin string) (string, error) {
	CoinTicker := strings.ToUpper(coin) + "USDT"
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + CoinTicker)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ticker Ticker
	err = json.NewDecoder(resp.Body).Decode(&ticker)
	if err != nil {
		return "", err
	}

	return ticker.Price, nil
}

/*
func main() {
	// Get the time for 12:00 PM UTC
	start := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 0, 0, 0, time.UTC)

	// If the current time is before 12:00 PM UTC, set the start time to the previous day
	if time.Now().Before(start) {
		start = start.Add(-24 * time.Hour)
	}

	// Wait until 12:00 PM UTC
	time.Sleep(start.Sub(time.Now()))

	// Continuously retrieve the BTC price every hour
	for {
		price, err := GetBTCPrice()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("The current price of BTC on Binance is $%.2f\n", price)
		time.Sleep(time.Hour)
	}
}

*/
