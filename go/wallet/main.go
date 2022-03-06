package main

import "fmt"

func main() {
	var c conversionRate = Rupee
	c.convertToBase(150)
	fmt.Println(c)

	money := createDollar(2, Dollar)
	fmt.Println(*money)
}
