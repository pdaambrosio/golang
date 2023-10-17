package main

import "fmt"

// func clone (s []float64) []float64 { // This is not a generic function
func main() {
	testScores := []float64 {
		87.3,
		92.2,
		0.0,
		-12.3,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c)
}

// This is not a generic function
func clone (s []float64) []float64 {
	result := make([]float64, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}
