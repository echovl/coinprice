package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const binanceURL = "https://api.binance.com/api/v3"
const timeout = 20 // timeout in seconds
const pricePlaceholder = "---"

type PriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	client := &http.Client{}
	printer := message.NewPrinter(language.English)

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	if len(os.Args) != 2 {
		fmt.Println("USAGE: coinprice [SYMBOL]\nEXAMPLE:\n\tcoinprice BTCUSDT")
		os.Exit(1)
	}

	symbol := os.Args[1]
	url := fmt.Sprintf("%s/ticker/price?symbol=%s", binanceURL, symbol)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		printPlaceholderAndExit()
	}

	resp, err := client.Do(req)
	if err != nil {
		printPlaceholderAndExit()
	}
	defer resp.Body.Close()

	var ticker PriceTicker
	if err = json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		printPlaceholderAndExit()
	}

	symbolPrice, err := strconv.ParseFloat(ticker.Price, 32)
	if err != nil {
		printPlaceholderAndExit()
	}

	printer.Printf("%.2f", symbolPrice)
}

func printPlaceholderAndExit() {
	fmt.Print(pricePlaceholder)
	os.Exit(1)
}
