package main

import (
	"fmt"
)

// // provide a comment about the type
// type User struct {
// 	ID    int
// 	FName string
// 	LName string
// 	Email string
// }

// // provide a comment
// type Group struct {
// 	role          string
// 	users         []User
// 	latestUser    User
// 	spaceAvailble bool
// }

func describeUser(u User) string {
	u.ID = 25
	desc := fmt.Sprintf("Name: %s %s, Id: %d, Email: %s", u.FName, u.LName, u.ID, u.Email)
	return desc
}

func describeGroup(group Group) string {
	var count int = len(group.users)
	var latestUser string = describeUser(group.lastUser)

	desc := fmt.Sprintf("This group %s has %d users. The latest user is %s. Has space available: %t", group.role, count, latestUser, group.spaceAvailable)
	return desc
}

func main7() {

	var users []User = make([]User, 5)

	for i := 0; i < 5; i++ {
		index := fmt.Sprint(i)
		users[i] = User{ID: i, FName: "First" + index, LName: "Last" + index, Email: "email" + index}
	}

	user := User{ID: 1, FName: "first", LName: "lst", Email: "email"}

	fmt.Println(describeUser(user))
	fmt.Println(user)

	group := Group{role: "users", users: users, lastUser: user}

	fmt.Println(describeGroup(group))

}
