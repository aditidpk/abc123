package main

import "fmt"

func main() {
	var c conversionRate = Rupee
	c.convertToBase(150)
	fmt.Println(c)

	wallet := &Wallet{balance: *(createRupee(500, Rupee))}
	fmt.Println(*wallet)

	money := createDollar(2, Dollar)
	fmt.Println(*money)

	wallet.deposit(*money)
	fmt.Println(*wallet)
}
