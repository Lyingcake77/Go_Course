package user

import (
	"errors"
	"fmt"
	"time"
)

type str string

func (text str) log() {
	fmt.Print(text)
}

type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

type Admin struct {
	email     string
	passsword string
	User
}

func NewAdmin(email, password string) Admin {
	var test str = "asd"
	test.log()
	return Admin{
		email:     email,
		passsword: password,
		User: User{
			firstName:   "",
			lastName:    "",
			dateOfBirth: "",
			createdAt:   time.Now(),
		},
	}
}
func New(firstname, lastname, dob string) (*User, error) {
	if firstname == "" || lastname == "" {
		return nil, errors.New("required fields midding")
	}
	return &User{
		firstName:   firstname,
		lastName:    lastname,
		dateOfBirth: dob,
		createdAt:   time.Now(),
	}, nil
}

func (appUser User) OutputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.dateOfBirth, appUser.createdAt)
}

//asterisk required to p[ass by ref and update the values.]
func (appUser *User) ClearUserName() {
	//shortcut for (*appUser).firstname
	appUser.firstName = ""
	appUser.lastName = ""
}
