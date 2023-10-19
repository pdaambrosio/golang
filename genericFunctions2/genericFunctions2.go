package main

import "fmt"

// Function sum calculates the sum of all values in a map if values are integers
func sumInt(m map[string]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}

// Function sumFloat calculates the sum of all values in a map if values are floats
func sumFloat(m map[string]float64) float64 {
	sum := 0.0
	for _, v := range m {
		sum += v
	}
	return sum
}

/*
Generic functions are functions that can be used with different types of arguments.
In Go, generic functions are supported. However, they are not implemented in the same way as in other languages.
*/

// Function sumGeneric calculates the sum of all values in a map if values are integers or floats
func sumGeneric[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// Another way to write the function sumGeneric

// Define an interface Number that can be either an int, int64, float32 or float64 (constraints)
type Number interface {
	~int | ~int64 | ~float32 | ~float64 // We need add "~" before the type to indicate that it is an interface type and not a concrete type, with this we can use type Number2 in the function sumGeneric2
}

// Function sumGeneric2 calculates the sum of all values in a map if values are integers or floats
func sumGeneric2[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// There are aproximation we can do to in the types of the arguments of the function

type Number2 int

func main() {
	// Create a map with string keys and integer values
	arr1 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		//"key4": 4.2, // This value is not an integer, so it will be generated an error (cannot use 4.2 (untyped float constant) as int value in map literal (truncated))
	}

	// Call function sumInt with map arr1 as argument
	fmt.Println("Integer function:")
	fmt.Println(sumInt(arr1))

	// Create a map with string keys and float values
	arr2 := map[string]float64{
		"key1": 1.1,
		"key2": 2.2,
		"key3": 3.3,
		//"key4": 4, // This value is not a float, so it will be generated an error (cannot use 4 (untyped int constant) as float64 value in map literal)
	}

	// Call function sumFloat with map arr2 as argument
	fmt.Println("\nFloat function:")
	fmt.Println(sumFloat(arr2))

	// Call function sumGeneric with maps arr1 and arr2 as argument
	fmt.Println("\nGeneric function:")
	fmt.Println(sumGeneric(arr1))
	fmt.Println(sumGeneric(arr2))

	// Call function sumGeneric2 with maps arr1 and arr2 as argument
	fmt.Println("\nGeneric function 2:")
	fmt.Println(sumGeneric2(arr1))
	fmt.Println(sumGeneric2(arr2))

	// Implementing type Number2
	var a, b, c Number2
	a = 1
	b = 2
	c = 3

	// Create a map with string keys and Number2 values
	arr3 := map[string]Number2{
		"key1": a,
		"key2": b,
		"key3": c,
	}

	// Call function sumGeneric2 with map arr3 as argument
	fmt.Println("\nGeneric function 2 with Number2:")
	fmt.Println(sumGeneric2(arr3)) // Its is possible because the "~" before the type Number that indicates that it is an interface type and not a concrete type
}
