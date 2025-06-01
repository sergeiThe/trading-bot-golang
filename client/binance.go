package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"trading-bot/models"
)

type Client struct {
	ApiUrl    string // TODO: Configure
	ApiKey    string
	ApiSecret string
	HeaderKey string
}

func (c *Client) FetchData() ([]models.Candle, error) {
	// return crypto.createHmac("sha256", secret).update(query).digest("hex");
	// const query = `symbol=${symbol}&interval=${interval}&limit=${limit}`;
	// [1748721600000,"104759.99000000","104852.94000000","104587.56000000","104834.02000000","94.23678000",1748725199999,"9868701.04692010",5847,"45.82295000","4800587.52146740","0"]
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
			Open:  parseStrToFloat(kline[1]),
			High:  parseStrToFloat(kline[2]),
			Low:   parseStrToFloat(kline[3]),
			Close: parseStrToFloat(kline[4]),
		}
		candles = append(candles, candle)
	}
	return candles
}


func parseStrToFloat(v any) float64 {
	f, _ := strconv.ParseFloat(v.(string), 64)
	return f
}
