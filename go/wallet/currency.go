package main

type conversionRate float64

const (
	Rupee  conversionRate = 1
	Dollar conversionRate = 74.85
)

func (cr conversionRate) convertToBase(amount float64) float64 {
	return amount * float64(cr)
}

func (cr conversionRate) convertToPreferredCurrency(amount float64, crt conversionRate) float64 {
	return cr.convertToBase(amount) / float64(crt)
}