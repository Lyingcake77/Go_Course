package main

import (
	"fmt"
)

type Product struct {
	title string
	id    int
	price float64
}

type floatMap map[string]float64

func main() {
	//1
	hobbies := [3]string{"thing1", "thing2", "thibng 3"}
	//2
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	hobbiesShort := hobbies[1:2]
	fmt.Println(hobbiesShort)
	//3
	hobbiesShort2 := hobbies[:2]
	fmt.Println(hobbiesShort2)

	//4
	//?

	fmt.Println(cap(hobbiesShort2))
	magicSlice := hobbiesShort2[1:3]
	fmt.Println(magicSlice)

	//5
	courseGoals := []string{"firstGoal", "secondGoal"}
	fmt.Println(courseGoals)
	courseGoals[1] = "newSecondGoal"
	fmt.Println(courseGoals)
	courseGoalsAdded := append(courseGoals, "third ghoul")
	fmt.Println(courseGoalsAdded)
	//6
	products := []Product{{title: "test2", id: 1, price: 20.6}, {title: "test3", id: 12, price: 450.6}}
	fmt.Println(products)

	newProduct := Product{title: "test4", id: 3, price: 64.6}
	fmt.Println(append(products, newProduct))

	test()
	mapStuff()
	makeingStuff()
}

func test() {
	// var test []float64

	prices := [4]float64{2, 3, 4.8}
	fmt.Println(prices)
	fmt.Println(prices[1])
	//slice like python, no negative slices
	fmt.Println(prices[1:2])
	fmt.Println(prices[:2])
	fmt.Println(prices[1:])
	fmt.Println(prices[0:2][:1])

	//all sub slices are part of the same "memory" object, bellow is a bastard way to change the value of prices
	temp := prices[1:2]
	temp[0] = 99
	fmt.Println(prices)

	//len shows the slice length, cap shows the memory value length, but only to the right of the array.
	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(temp), cap(temp))

	//dynamic array
	dynamicArray := []float64{12.22, 4}
	fmt.Println(dynamicArray)
	//append creates a new array with new memory when creating a larger array than initially set
	newArray := append(dynamicArray, 5.77)
	newArray = append(newArray, 5.77, 33.4)
	//this can decomstruct the array into indiviidual items
	newArray = append(newArray, []float64{5.77, 123.23, 1234}...)
	fmt.Println(newArray)

}
func mapStuff() {
	//maps [key] value
	websites := map[string]string{"google": "google.com", "aws": "aws.com"}
	fmt.Println(websites)
	websites["aws"] = "https//aws.com"
	fmt.Println(websites)
	websites["Newkey"] = "linkin.com"
	delete(websites, "google")
	fmt.Println(websites)
}

func makeingStuff() {
	//creates an initial array of length 10, but still lets you use append.
	//also makes a limit of 100 records
	usernames := make([]string, 10, 100)
	usernames[9] = "asd"
	usernames = append(usernames, "Max")
	fmt.Println(usernames)

	//preallocating memory for maps
	courses := make(map[string]float64, 2)
	courses["go"] = 4.5
	fmt.Println(courses)
	//typeAlias
	courses = make(floatMap, 2)
	courses["go"] = 4.5
	courses["react"] = 4.6
	fmt.Println(courses)

	for index, value := range usernames {
		fmt.Println(index, value)
	}
	for key, value := range courses {
		fmt.Println(key, value)
	}

}
