package main

import (
	"fmt"
)

func main_old() {
	var age int = 29

	var fage float64 = 29

	// var age = 32

	var name1, name2 = "one", "test"

	name := "hello"

	fmt.Println("Canonical hello world")

	fmt.Printf("This is my name: %s, number is %d %.2f %s %s %s", "Seattle", age, fage, name1, name2, name)

	value, err := fmt.Println("Testing..")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(value)
	}

	if val, err := fmt.Println("Say what"); err == nil {
		fmt.Println("Here we go", val)
	}
}
