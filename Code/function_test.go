package main

import "testing"

func TestAverage(t *testing.T) {
	expected := 6.0
	actual := getAverage([]float64{2.5, 5, 7.5, 10, 5})

	if actual != expected {
		t.Errorf("Incorrect expected: %f actual: %f", expected, actual)
	}
}
