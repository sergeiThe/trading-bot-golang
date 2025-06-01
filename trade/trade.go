package trade


/*
#cgo LDFLAGS: -lta-lib -lm
*/
import "C"

import "github.com/d4l3k/talib"


func Run() {
	values := []float64{1,2,3}
	talib.Acos(values)
}




