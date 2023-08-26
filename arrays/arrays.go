package main

import "fmt"

/*
Arrays are fixed size, and cannot be resized.
Arrays are zero indexed.
Arrays must be of the same type.
*/

func main() {
	var array [5]int // array of 5 integers

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println()

	// Short hand declaration
	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	fmt.Println(array2[0])
}
