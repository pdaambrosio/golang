package main

import "fmt"

func learningSlice() {
	// Array are fixed length
	// Array are comparable if the type of the array is comparable
	names := [8]string{"John", "Wick", "Ethan", "Hunt", "Jason", "Bourne", "James", "Bond"} // length 8, capacity 8

	// Print the array
	fmt.Println(names, len(names), cap(names))
	
	// Print the length of the array
	fmt.Println(len(names))

	// Print the capacity of the array
	fmt.Println(cap(names))

	// Append the array
	names = append(names, "Jack", "Reacher")

	// Slice are dynamic length
	fmt.Println(names, len(names), cap(names))
	fmt.Println()
}

func newSliceLengthCapacity() {
	// Slice declaration
	// Slice are dynamic length
	// Slice aren't comparable
	names := make([]string, 5, 10) // length 5, capacity 10
	fmt.Println(names, len(names), cap(names))
	fmt.Println()
}

func learningMap() {
	// Maps are key value pairs
	// Maps are unordered list
	ages := make(map[string]uint8) // Map declaration
	ages["John"] = 30 // Add key value pair
	ages["Ethan"] = 50
	ages["Jason"] = 40

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

func learningStruct() {
	// Struct is a collection of fields
	// Private field is not printed because it is not exported
	// Private filed is not accessible outside the package
	type Person struct {
		Name string
		LastName string
		Age uint8
		Status bool
		cpf string // Private field
	}

	// Struct declaration without initialization
	var p1 Person
	p1.Name = "John"
	p1.LastName = "Wick"
	p1.Age = 30
	p1.Status = true
	p1.cpf = "123.456.789-00" // Private field

	// Struct declaration with initialization
	p2 := Person {
		Name: "Ethan",
		LastName: "Hunt",
		Age: 50,
		Status: true,
		cpf: "987.654.321-00",
	}

	// Struct declaration with initialization without field name
	var p3 Person = Person {
		"Jason",
		"Bourne",
		40,
		true,
		"456.789.123-00",
	}

	// Print the structs
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

}

func main() {
	learningSlice()
	newSliceLengthCapacity()
	learningMap()
	learningStruct()
}
