package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trading-bot/models"
	"trading-bot/utils"
)


type Binance struct {
	ApiUrl    string // TODO: Configure
	ApiKey    string
	ApiSecret string
	HeaderKey string
}

func (c *Binance) FetchData() ([]models.Candle, error) {
	// TODO: This should come from outside
	symbol := "BTCUSDC"
	interval := "1h"
	limit := 1000

	path := "/api/v3/klines"
	query := fmt.Sprintf("symbol=%s&interval=%s&limit=%d", symbol, interval, limit)
	url := c.ApiUrl + path + "?" + query

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(c.HeaderKey, c.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	var klines [][]any

	err = json.Unmarshal(bytes, &klines)
	if err != nil {
		return nil, err
	}

	candles := parseKlines(klines)

	return candles, nil

}

func parseKlines(klines [][]any) []models.Candle {
	candles := []models.Candle{}

	for i := range klines {
		kline := klines[i]
		candle := models.Candle{
			Open:  utils.ParseStrToFloat(kline[1]),
			High:  utils.ParseStrToFloat(kline[2]),
			Low:   utils.ParseStrToFloat(kline[3]),
			Close: utils.ParseStrToFloat(kline[4]),
		}
		candles = append(candles, candle)
	}
	return candles
}


