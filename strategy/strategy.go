package strategy

import (
	"errors"
	"log"
	"trading-bot/models"
)

type Signal struct {
	Action string
	Reason string
}

const (
	BUY = "BUY"
	SELL = "SELL"
	HOLD = "HOLD"
)

func GenSignal(candles []models.Candle) (*Signal, error) {

	if len(candles) < 15 {
		return nil, errors.New("not enough data for RSI")
	}

	rsi := calculateRSI(candles, 14)
	latestRSI := rsi[len(rsi)-1]
	log.Printf("RSI: %f \n", latestRSI)

	switch {
	case latestRSI < 30:
		return &Signal{
			Action: BUY,
			Reason: "RSI below 30 (oversold)",
		}, nil
	case latestRSI > 70:
		return &Signal{
			Action: SELL,
			Reason: "RSI above 70 (overbought)",
		}, nil
	default:
		return &Signal{
			Action: HOLD,
			Reason: "RSI in the middle",
		}, nil // no signal
	}
}

func calculateRSI(candles []models.Candle, period int) []float64 {
	if len(candles) <= period {
		return []float64{}
	}

	var gains, losses []float64
	for i := 1; i <= period; i++ {
		change := candles[i].Close - candles[i-1].Close
		if change >= 0 {
			gains = append(gains, change)
			losses = append(losses, 0)
		} else {
			gains = append(gains, 0)
			losses = append(losses, -change)
		}
	}

	avgGain := sum(gains) / float64(period)
	avgLoss := sum(losses) / float64(period)

	rsis := make([]float64, 0, len(candles)-period)
	rs := avgGain / avgLoss
	rsis = append(rsis, 100-(100/(1+rs)))

	for i := period + 1; i < len(candles); i++ {
		change := candles[i].Close - candles[i-1].Close
		gain := 0.0
		loss := 0.0

		if change >= 0 {
			gain = change
		} else {
			loss = -change
		}

		avgGain = (avgGain*float64(period-1) + gain) / float64(period)
		avgLoss = (avgLoss*float64(period-1) + loss) / float64(period)

		rs := avgGain / avgLoss
		rsi := 100 - (100 / (1 + rs))
		rsis = append(rsis, rsi)
	}

	return rsis
}

func sum(data []float64) float64 {
	total := 0.0
	for _, v := range data {
		total += v
	}
	return total
}
