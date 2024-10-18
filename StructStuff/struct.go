package main

import (
	"example.com/structure-stuff/user"
	"fmt"
)

func main() {

	firstname := getUserData("firstName")
	lastname := getUserData("last name")
	dob := getUserData("DOB")

	var appUser *user.User
	// var err error
	appUser, err := user.New(firstname, lastname, dob)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
	// outputUserDetailPointer(appUser)
	appUser.ClearUserName()
	appUser.OutputUserDetails()
	
	admin := user.NewAdmin("email@email.com", "1234567")
	admin.OutputUserDetails()
}

// func outputUserDetailPointer(appUser *user.User) {
// 	fmt.Println(appUser.firstName, appUser.lastName, appUser.dateOfBirth, appUser.createdAt)
// }

//True horror
func getUserData(PromptText string) string {
	fmt.Print(PromptText)
	var value string
	fmt.Scanln(&value)
	return value
}
