package simplenum

import (
	"github.com/shopspring/decimal"
)

//四捨五入
func Round(v float64, decimals int) float64 {
	dv := decimal.NewFromFloat(v)
	dv = dv.Round(int32(decimals))
	a, _ := dv.Float64()
	return a
}

//無條件進位
func Ceil(v float64, decimals int) float64 {
	pow := decimal.NewFromFloat(10)
	pow = pow.Pow(decimal.NewFromFloat(float64(decimals)))
	dv := decimal.NewFromFloat(v)
	ans := dv.Mul(pow)
	ans = ans.Ceil()
	ans = ans.Div(pow)
	a, _ := ans.Float64()
	return a
}

//無條件捨去
func Floor(v float64, decimals int) float64 {
	pow := decimal.NewFromFloat(10)
	pow = pow.Pow(decimal.NewFromFloat(float64(decimals)))
	dv := decimal.NewFromFloat(v)
	ans := dv.Mul(pow)
	ans = ans.Floor()
	ans = ans.Div(pow)
	a, _ := ans.Float64()
	return a
}
