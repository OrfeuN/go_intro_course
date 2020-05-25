package main

import (
	"fmt"
)

func main2() {
	sentence := "This is a sentence"

	for index, letter := range sentence {
		if index%2 == 1 {
			fmt.Print(string(letter))
		}
	}

	var name string
	var age int

	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println("Enter your age")
	fmt.Scan(&age)
	fmt.Println("Age", age, " Name:", name)

}
