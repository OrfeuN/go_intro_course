package main

import "fmt"

type Point struct {
	X, Y int
}

var c Point = Point{X: 5, Y: 25}

func main10() {
	var name string = "Myname"
	var namePointer *string = &name
	var whatName string = *namePointer

	var item int = 25

	itemPointer := &item

	fmt.Println(name, namePointer, whatName, item, itemPointer, *itemPointer)

	point := c
	point.X = 25

	fmt.Println(point)
	fmt.Println(c)

	changePoint := &c
	changePoint.X = 30 // go allows saying changePoint instead of having to say (*changePoint)
	(*changePoint).Y = 30

	fmt.Println(changePoint, *changePoint, c)
}
