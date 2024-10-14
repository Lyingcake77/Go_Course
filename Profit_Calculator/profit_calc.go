package main

import (
	"fmt"
	// "math"
)

func main() {

	revenue, expenses, taxRate := GetUserInputs()
	earningBeforeTax, earningAfterTax, ratio := CalculateEarnings(revenue, expenses, taxRate)

	formatedReturnText := fmt.Sprintf(`
	EBT: %.2f
	Profit: %.2f
	Ratio: %v
	`, earningBeforeTax, earningAfterTax, ratio)

	fmt.Println(formatedReturnText)
}

func CalculateEarnings(Revenue, Expenses, TaxRate float64) (EBT float64, Profit float64, Ratio float64) {
	EBT = Revenue - Expenses
	Profit = EBT * (1 - TaxRate/100)
	Ratio = EBT / Profit
	return EBT, Profit, Ratio
}

func GetUserInputs() (Revenue float64, Expsenses float64, TaxRate float64) {
	//get user input
	fmt.Print("enter revenue: ")
	fmt.Scan(&Revenue)
	fmt.Print("enter expenses: ")
	fmt.Scan(&Expsenses)
	fmt.Print("enter taxRate: ")
	fmt.Scan(&TaxRate)
	return
}
