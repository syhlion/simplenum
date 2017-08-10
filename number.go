package simplenum

import "math"

//四捨五入
func Round(v float64, decimals int) float64 {
	var pow float64 = 1

	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(int((v*pow)+0.5)) / pow
}

//無條件進位
func Ceil(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(math.Ceil(v*pow)) / pow
}

//無條件捨去
func Floor(v float64, decimals int) float64 {
	var pow float64 = 1
	for i := 0; i < decimals; i++ {
		pow *= 10
	}
	return float64(math.Floor(v*pow)) / pow
}
