package main

import (
	"coinconv/clients/coin_market"
	"coinconv/internal/coin_conv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	market := coin_market.NewCoinMarket()
	conv := coin_conv.NewCoinConv(&market)
	consoleArgs := os.Args
	if len(consoleArgs) < 4 {
		fmt.Printf("Error: %v", "wrong count of arguments")
	}

	count, err := strconv.ParseFloat(consoleArgs[1], 64)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}

	convertedData, err := conv.ConvertData(count, consoleArgs[3], consoleArgs[2])
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}
	fmt.Println(convertedData)
}
