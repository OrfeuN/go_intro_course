package main

import "fmt"

// can be exported; needs comment
type Describer interface {
	describe() string
}

type User struct {
	ID                  int
	FName, LName, Email string
}

type Group struct {
	role           string
	users          []User
	lastUser       User
	spaceAvailable bool
}

func (u User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Id: %d, Email: %s", u.FName, u.LName, u.ID, u.Email)
	return desc
}

func (group Group) describe() string {
	var count int = len(group.users)
	var latestUser string = group.lastUser.describe()

	desc := fmt.Sprintf("This group %s has %d users. The latest user is %s. Has space available: %t", group.role, count, latestUser, group.spaceAvailable)
	return desc
}

// describe struct passed in
func Describing(d Describer) string {
	return d.describe()
}

func mainInt() {
	var users []User = make([]User, 5)

	for i := 0; i < 5; i++ {
		index := fmt.Sprint(i)
		users[i] = User{ID: i, FName: "First" + index, LName: "Last" + index, Email: "email" + index}
	}

	user := User{ID: 1, FName: "first", LName: "lst", Email: "email"}

	fmt.Println(user.describe())
	fmt.Println(user)

	group := Group{role: "users", users: users, lastUser: user}

	fmt.Println(group.describe())
	fmt.Println(group)

	fmt.Println(" descring", Describing(user))

	fmt.Println(" descring", Describing(group))
}
