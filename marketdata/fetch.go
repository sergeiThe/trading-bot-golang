package marketdata

import (
	m "trading-bot/models"
)


type Candle m.Candle

func Format(open, low, high, close []float64) []Candle {

	candles := []Candle{}

	for i := range open {
		candles = append(candles, Candle{
			Open: open[i],
			Close: close[i],
			Low: low[i],
			High: high[i],
		})
	}

	return []Candle{}
} 
