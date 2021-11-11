package main

import (
	"coinconv/clients/coin_market"
	"coinconv/internal/coin_comission"
	"coinconv/internal/coin_conv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	market, err := coin_market.NewCoinMarket()
	if err != nil {
		fmt.Printf("Can't create coin market client: %v", err)
		return
	}
	conv := coin_conv.NewCoinConv(market)
	consoleArgs := os.Args
	if len(consoleArgs) < 4 {
		fmt.Printf("Error: %v", "wrong count of arguments")
		return
	}

	count, err := strconv.ParseFloat(consoleArgs[1], 64)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	commission := 0.1

	convertedData, err := conv.ConvertData(count, consoleArgs[3], consoleArgs[2])
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}

	commissionCounter, err := coin_comission.NewCommissionCalculator(conv, commission)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	commissionPrice, err := commissionCounter.ConvertData(count, consoleArgs[3], consoleArgs[2])
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	fmt.Println(convertedData)
	fmt.Println(commissionPrice)
}
