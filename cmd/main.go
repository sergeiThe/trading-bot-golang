package main

import (
	"fmt"
	"log"
	"os"
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
	discordWH := os.Getenv("DISCORD_WEBHOOK")

	c := client.Client{
		ApiUrl:    apiUrl,
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		HeaderKey: headerKey,
	}

	candles, err := c.FetchData()
	if err != nil {
		log.Fatalf("Fetching data failed: %v", err)
	}

	signal, _ := strategy.GenSignal(candles)

	log.Printf("SIGNAL: %v\n", signal)

	webhookUrl := discordWH
	err = notification.SendDiscordNotification(webhookUrl, signal.Reason)
	if err != nil {
		log.Println("Failed to send Discord notification:", err)
	}
}
