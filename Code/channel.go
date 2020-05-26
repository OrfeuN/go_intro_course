package main

import (
	"fmt"
)

func add(val chan int, items ...int) {
	sum := 0
	for _, num := range items {
		sum += num
	}
	val <- sum
}

func main() {
	var c chan int = make(chan int)

	go add(c, -20, 20, 10, 30, 20, 20, 30)

	go add(c, 10, 20, 30, 30, 20, 20, 30)

	sum1, sum2 := <-c, <-c

	fmt.Println(sum1, sum2, sum1+sum2)
}
