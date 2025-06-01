package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"trading-bot/client"
	"trading-bot/notification"
	"trading-bot/strategy"

	dotenv "github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start")
	err := dotenv.Load()
	if err != nil {
		panic("Could not load env")
	}

	// TODO: Check empty values
	apiKey := os.Getenv("API_KEY")
	apiUrl := os.Getenv("API_URL")
	apiSecret := os.Getenv("API_SECRET")
	headerKey := os.Getenv("HEADER_KEY")

	notifProviders := notification.InitProviders()
	strategies := strategy.InitStrategies()

	c := client.Binance{
		ApiUrl:    apiUrl,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		HeaderKey: headerKey,
	}

	for {
		candles, err := c.FetchData()

		if err != nil {
			log.Fatalf("Fetching data failed: %v", err)
		}

		signals := strategy.Run(strategies, candles)
		strategy.ProcessSignals(signals, func(s strategy.Signal) {
			if s.Action == strategy.BUY {
				errors := notification.Run(s.Reason, notifProviders)
				notification.ProcessErrors(errors)
			}
		})

		time.Sleep(time.Second * 5)
	}
}
