package main

import "fmt"

/*
Slices are like arrays, but their size is dynamic.
Slices are zero indexed.
Slices must be of the same type.
*/

func main() {
	// var slice []int // slice of integers
	slice := []int{1, 2, 3} // slice of 3 integers
	fmt.Println(slice)

	s1 := slice[1:] // slice from index 1 to the end
	fmt.Println(s1)

	s2 := slice[:2] // slice from the beginning to index 2
	fmt.Println(s2)

	s3 := slice[1:2] // slice from index 1 to index 2
	fmt.Println(s3)
}
