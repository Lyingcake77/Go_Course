package main

import (
	"example.com/banking/fileManagement"
	"fmt"
)

const AccountBalanceFile string = "balance.txt"

var Balance float64 = 0

func main() {
	Balance, err := fileManagement.GetFloatFromFile(AccountBalanceFile)
	if err != nil {
		fmt.Print(err)
		//panic("cant continue")
	}
	fmt.Println("Welcome to Bank")
	for {
		choice := GetUserChoiceInputs()
		if choice == 1 {
			//print balance
			PrintBalance(Balance)
		} else if choice == 2 {
			//deposit money
			money := GetUserMoneyInputs("How much would you like to Deposit")
			Balance += money
			fileManagement.SaveFloatToFile(AccountBalanceFile, Balance)
			PrintBalance(Balance)
		} else if choice == 3 {
			//withdraw money
			money := GetUserMoneyInputs("How much would you like to Withdraw")
			if money <= Balance {
				Balance -= money
				fileManagement.SaveFloatToFile(AccountBalanceFile, Balance)
			} else {
				fmt.Print("Not Enough Funds")
			}
			PrintBalance(Balance)
		} else if choice == 4 {
			//exist banking system
			fmt.Println("Thank you for Banking with Go Inc.")
			break
		} else {
			fmt.Println("error, values not between 1 and 4!!")
		}
	}

}

func GetUserMoneyInputs(DisplayText string) (Money float64) {
	fmt.Print(DisplayText)
	fmt.Scan(&Money)
	return Money
}
