package main

import (
	"fmt"
)

func main6() {
	var nums []float64 = make([]float64, 5, 10)

	nums[0] = 2.5
	nums[1] = 5
	nums[2] = 7.5
	nums[3] = 10
	nums[4] = 12.5

	items := append(nums, 15)

	var average = getAverage(items)

	fmt.Println("Avergae", average)

	pets := map[string]string{"fido": "dog", "jack": "dog", "kitty": "cat"}
	fmt.Println(pets)
	fmt.Println("Has pet", hasItem(pets, "fido"))
	fmt.Println("Has pet", hasItem(pets, "nano"))

}

func hasItem(pets map[string]string, key string) bool {
	_, ok := pets[key]
	return ok
}

func getAverage(items []float64) float64 {
	sum := 0.0
	for _, value := range items {
		sum += value
	}
	if len(items) == 0 {
		return 0.0
	}
	return sum / float64(len(items))
}
