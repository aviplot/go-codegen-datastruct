package main


import (
	"github.com/aviplot/go-finance-math/financial"
	"github.com/shopspring/decimal"
	"github.com/aviplot/go-mortgage-calc/mortgage"
)
func main() {
	d := financial.NewDateFromFormattedString("2020-05-20")
	a := decimal.NewFromFloat(1.23)
	mortgage.GetShpitzerCashflowInstance()
}



