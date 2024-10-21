package main

import (
	"fmt"
)

type TransformFn func(int) int

func main() {
	dNumbers := []int{1, 2, 3}
	fmt.Println(dNumbers)
	doubled := transformNumbers(&dNumbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&dNumbers, func(num int) int { return num * 3 })
	fmt.Println(tripled)

	quad := transformNumbers(&dNumbers, createTransformer(4))
	fmt.Println(quad)
	test := getTransformerFunction()
	fmt.Println(test)

	fmt.Println(factorial(5))

	fmt.Println(sumup([]int{1, 2, 3}))

	fmt.Println(sumup2(1, 2, 3))
	fmt.Println(sumup2([]int{1, 2, 3}...))
}

// input a function as a value
func transformNumbers(numbers *[]int, transform TransformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(a int) int {
	return a * 2
}

// return function as a function
func getTransformerFunction() TransformFn {
	return double
}

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		//scope applies like an if statement here.
		//this will also stay scoped to this call and it does not affect any other call that
		//uses this function with different values.
		return num * factor
	}
}

// dont do this in a day to day usecase.
func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

// variadic function?>
func sumup(num []int) int {
	sum := 0
	for _, val := range num {
		sum += val
	}
	return sum
}
func sumup2(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
