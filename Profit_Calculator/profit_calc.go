package main

import (
	"fmt"
	// "math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	//get user input
	fmt.Print("enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("enter taxRate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax

	fmt.Println("EBT:", earningBeforeTax)
	fmt.Println("Profit:", earningAfterTax)
	fmt.Println("ratio:", ratio)
}
