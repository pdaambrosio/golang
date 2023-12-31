package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3.14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

// addable is a custom type that can be used as a constraint
type addable interface {
	int | float64 | string
}

// add sums the elements of a slice. The elements must be of a type
// that implements the addable interface
func add[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}

	return result
}
