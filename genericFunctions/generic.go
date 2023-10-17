package main

import "fmt"

// func clone (s []float64) []float64 { // This is not a generic function
func main() {
	testScores := []float32 { // if I change this to a slice of int, the clone function will not work
		87.3,
		92.2,
		0.0,
		-12.3,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c)
}

// This is not a generic function
/*
func clone (s []float64) []float64 { // Its type is []float64 -> it only works with float64
	result := make([]float64, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}
*/

// This is a generic function, it works with any type of slice (int, float, string, etc) -> []V (V is a type parameter)
func clone[V any] (s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}