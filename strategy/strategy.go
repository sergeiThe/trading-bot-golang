package strategy

import (
	"log"
	"trading-bot/models"
)

type Signal struct {
	Action string
	Reason string
}

const (
	BUY  = "BUY"
	SELL = "SELL"
	HOLD = "HOLD"
)

type Strategy interface {
	GenSignal(candles []models.Candle) (*Signal, error)
}

func InitStrategies() []Strategy {
	return []Strategy{
		RSIStrategy{},
	}
}

func Run(strategies []Strategy, candles []models.Candle) []Signal {

	signals := []Signal{}
	for _, str := range strategies {
		signal, err := str.GenSignal(candles)
		if err != nil {
			log.Fatalf("Signal generating error: %v\n", err)
		}
		signals = append(signals, *signal)
	}
	return signals
}

func ProcessSignals(signals []Signal, operation func(Signal)) {
	for _, sig := range signals {
		operation(sig)
	}
}
