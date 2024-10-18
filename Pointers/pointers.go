package main

import (
	"fmt"
)

func main() {
	age := 32
	agePointer := &age

	fmt.Println("age", age)
	adultYears := getAdultYears(agePointer)
	getAdultYearsPointer(agePointer)
	fmt.Println(adultYears)
	fmt.Println(age)
}

//True horror
func getAdultYearsPointer(age *int) {
	*age = *age - 18
}

func getAdultYears(age *int) int {
	return *age - 18
}
