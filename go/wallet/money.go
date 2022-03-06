package main

type Money struct {
	amount   float64
	currency conversionRate
}

func createRupee(amount float64, currency conversionRate) *Money {
	money := &Money{amount: amount, currency: currency}
	return money
}

func createDollar(amount float64, currency conversionRate) *Money {
	money := &Money{amount: amount, currency: currency}
	return money
}
