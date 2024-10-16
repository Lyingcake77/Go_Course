package main

import "fmt"

func GetUserChoiceInputs() (Choice int) {
	fmt.Println(`What do you want to do?
	1. Check balace
	2. Deposit money
	3. Withdraw money
	4. Exit`)
	fmt.Scan(&Choice)
	return Choice
}
func PrintBalance(Balance float64) {
	fmt.Println("Current Balance: $", Balance)
}
