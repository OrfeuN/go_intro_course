package main

import "fmt"

func average(values ...float64) float64 {
	// sum, count := 0.0, 0.0
	// for _, value := range values {
	// 	sum += value
	// 	count++
	// }
	// if count == 0 {
	// 	return 0
	// }

	// return sum / count
	sum := 0.0

	for _, value := range values {
		sum += value
	}

	if len(values) == 0 {
		return 0
	}
	var count float64 = float64(len(values))
	return sum / count
}

func main3() {
	fmt.Println(average(3.0, 3, 5, 9.9))
}
