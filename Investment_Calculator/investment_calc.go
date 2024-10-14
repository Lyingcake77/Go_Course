package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	//recomended appreach for assigning infered values
	expectedReturnRate := 5.5
	var yearsToHoldInvestment float64

	//get user input
	fmt.Print("enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("enter return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("enter years: ")
	fmt.Scan(&yearsToHoldInvestment)
	futureValue, futureRealValue := CalculateInvestentValues(investmentAmount, expectedReturnRate, yearsToHoldInvestment)

	fmt.Println("future Value:", futureValue)
	//multiline formatted formatted text
	formatResult := fmt.Sprintf(
		`Real Future Value: %.2f
multiLine`, futureRealValue)
	fmt.Println(formatResult)
}

func CalculateInvestentValues(investmentAmount, expectedReturnRate, yearsToHoldInvestment float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, yearsToHoldInvestment)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, yearsToHoldInvestment)
	//this returns the specified values. I also dont like this shorthand and instead should return the values explicitly
	return
	// return futureValue, futureRealValue
}
