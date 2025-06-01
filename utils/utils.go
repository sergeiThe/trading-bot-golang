package utils

import "strconv"





func ParseStrToFloat(v any) float64 {
	f, _ := strconv.ParseFloat(v.(string), 64)
	return f
}
