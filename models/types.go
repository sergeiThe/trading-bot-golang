package models

type Candle struct {
	Open  float64 `json:"open"`
	Close float64 `json:"close"`
	Low   float64 `json:"low"`
	High  float64 `json:"high"`
}
