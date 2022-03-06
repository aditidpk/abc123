package main

type Money struct {
	amount   float64
	currency conversionRate
}

func (m Money) add(money Money) *Money {
	summedAmount := m.currency.convertToPreferredCurrency(m.amount, money.currency) + money.amount
	newMoney := Money{amount: summedAmount, currency: money.currency}
	return &newMoney
}

func (m Money) subtract(money Money) *Money {
	subtractedAmount := m.currency.convertToPreferredCurrency(m.amount, money.currency) - money.amount
	newMoney := Money{amount: subtractedAmount, currency: money.currency}
	return &newMoney
}

func (m Money) convertToPreferredCurrency(preferredCurrency conversionRate) *Money {
	convertedAmount := m.currency.convertToPreferredCurrency(m.amount, preferredCurrency)
	newMoney := Money{amount: convertedAmount, currency: preferredCurrency}
	return &newMoney
}

func createRupee(amount float64, currency conversionRate) *Money {
	money := &Money{amount: amount, currency: currency}
	return money
}

func createDollar(amount float64, currency conversionRate) *Money {
	money := &Money{amount: amount, currency: currency}
	return money
}
