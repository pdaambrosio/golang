package main

import "fmt"

/*
Map is a collection of key value pairs
Map is unordered list
Map are reference type
Map are comparable if the type of the key is comparable
*/

func main() {
	// Map declaration
	// ages := map[string]uint8{} // Map declaration
	// ages := make(map[string]uint8) // Map declaration
	ages := map[string]uint8{ // Map declaration
		"John":  30,
		"Ethan": 50,
		"Jason": 40,
	}

	// Print the map
	fmt.Println(ages) // Unordered list
	fmt.Println(ages["John"]) // Print the value of the key
	fmt.Println(ages["Malfoy"]) // Print the value if the key is not present in the map return default value of the type

	// Check if the key is present in the map
	age1, ok1 := ages["Malfoy"] // ok1 is true if the key is present in the map
	age2, ok2 := ages["John"] // ok2 is true if the key is present in the map
	fmt.Println(age1, ok1)
	fmt.Println(age2, ok2)

	// Delete the key value pair
	delete(ages, "John")
	fmt.Println(ages)
	fmt.Println()
}
