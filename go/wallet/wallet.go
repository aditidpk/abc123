package main

type Wallet struct {
	balance Money
}

func (w *Wallet) deposit(money Money) {
	w.balance = *w.balance.add(money)
}

func (w *Wallet) withdraw(money *Money) {
	w.balance = *w.balance.subtract(*money)
}

func (w *Wallet) balanceOf(preferredCurrency conversionRate) *Money {
	return w.balance.convertToPreferredCurrency(preferredCurrency)
}
