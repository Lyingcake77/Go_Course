package main

import (
	"fmt"
	"os"
	"strconv"
)

const AccountBalanceFile string = "balance.txt"

var Balance float64 = 0

func main() {
	Balance = GetBalanceFromFile()
	fmt.Println("Welcome to Bank")
	for {
		choice := GetUserChoiceInputs()
		if choice == 1 {
			//print balance
			PrintBalance()
		} else if choice == 2 {
			//deposit money
			money := GetUserMoneyInputs("How much would you like to Deposit")
			Balance += money
			SaveBalanceToFile(Balance)
			PrintBalance()
		} else if choice == 3 {
			//withdraw money
			money := GetUserMoneyInputs("How much would you like to Withdraw")
			if money <= Balance {
				Balance -= money
				SaveBalanceToFile(Balance)
			} else {
				fmt.Print("Not Enough Funds")
			}
			PrintBalance()
		} else if choice == 4 {
			//exist banking system
			fmt.Println("Thank you for Banking with Go Inc.")
			break
		} else {
			fmt.Println("error, values not between 1 and 4!!")
		}
	}

}

func GetUserChoiceInputs() (Choice int) {
	fmt.Println(`What do you want to do?
	1. Check balace
	2. Deposit money
	3. Withdraw money
	4. Exit`)
	fmt.Scan(&Choice)
	return Choice
}
func GetUserMoneyInputs(DisplayText string) (Money float64) {
	fmt.Print(DisplayText)
	fmt.Scan(&Money)
	return Money
}
func PrintBalance() {
	fmt.Println("Current Balance: $", Balance)
}

func SaveBalanceToFile(Balance float64) {
	var balanceString string = fmt.Sprintf("%v", Balance)
	os.WriteFile(AccountBalanceFile, []byte(balanceString), 0644)
}

func GetBalanceFromFile() {
	data, _ := os.ReadFile(AccountBalanceFile)
	returnData, _ = strconv.ParseFloat(string(data), 64)
	return returnData
}
