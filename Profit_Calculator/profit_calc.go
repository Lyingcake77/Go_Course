package main

import (
	"errors"
	"fmt"
	"os"
	// "math"
)

const AccountBalanceFile string = "Earnings.txt"

func main() {

	revenue, expenses, taxRate, err  := GetUserInputs()
	if err != nil {
		fmt.Println(err)
		panic("there was an error getting the user inputs ")
	}
	earningBeforeTax, earningAfterTax, ratio := CalculateEarnings(revenue, expenses, taxRate)

	formatedReturnText := fmt.Sprintf(`
	EBT: %.2f
	Profit: %.2f
	Ratio: %v
	`, earningBeforeTax, earningAfterTax, ratio)
	SaveTextToFile(formatedReturnText)
	fmt.Println(formatedReturnText)
}

func CalculateEarnings(Revenue, Expenses, TaxRate float64) (EBT float64, Profit float64, Ratio float64) {
	EBT = Revenue - Expenses
	Profit = EBT * (1 - TaxRate/100)
	Ratio = EBT / Profit

	return EBT, Profit, Ratio
}

func GetUserInputs() (Revenue float64, Expsenses float64, TaxRate float64, Error error) {
	//get user input
	fmt.Print("enter revenue: ")
	fmt.Scan(&Revenue)
	fmt.Print("enter expenses: ")
	fmt.Scan(&Expsenses)
	fmt.Print("enter taxRate: ")
	fmt.Scan(&TaxRate)

	if Revenue <= 0 || Expsenses <= 0 || TaxRate <= 0 {
		return Revenue, Expsenses, TaxRate, errors.New("cannot accept 0 or negative number in user inputs")

	}
	return Revenue, Expsenses, TaxRate, nil
}

func SaveTextToFile(SaveText string) {
	os.WriteFile(AccountBalanceFile, []byte(SaveText), 0644)
}
