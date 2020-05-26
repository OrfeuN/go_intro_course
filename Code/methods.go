package main

import "fmt"

// type User struct {
// 	ID    int
// 	FName string
// 	LName string
// 	Email string
// }

// use this if state is to be modified
// func (u *User) describe() string {
// 	desc := fmt.Sprintf("Name: %s %s, Id: %d, Email: %s", u.FName, u.LName, u.ID, u.Email)
// 	return desc
// }

// func (u User) describe() string {
// 	desc := fmt.Sprintf("Name: %s %s, Id: %d, Email: %s", u.FName, u.LName, u.ID, u.Email)
// 	return desc
// }

func describe(u User) string {
	return u.describe()
}

func mainMethods() {
	user := User{ID: 1, FName: "first", LName: "lst", Email: "email"}

	fmt.Println(describe(user))
	fmt.Println(user.describe())
	fmt.Println(user)
}
