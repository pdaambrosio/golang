package main

import (
	"fmt"
)

type People struct {
	Name string
	LastName string
	Age uint8
	Status bool
	documentID string // Private field
}

type Category struct {
	Name string
	Father *Category // Pointer to Category struct
}

// Method is a function with a receiver argument
// The receiver appears in its own argument list between the func keyword and the method name
func (c Category) show() {
	fmt.Println(c.Name) // Print the struct field
}

func (c Category) HasFather() bool {
	return c.Father != nil // Return true if the father is not nil
}

func (p People) show() {
	fmt.Println(p.Name, p.LastName, p.
	Age, p.Status, p.documentID) // Print the struct fields
}

// This method is a value receiver
// It cannot modify the value to which the receiver points
func (p People) string() string {
	return fmt.Sprintf("%s %s: %d years old", p.Name, p.LastName, p.Age) // Return the string
}

// This method is a pointer receiver
// It can modify the value to which the receiver points
func (p *People) birthday() {
	p.Age++ // Increment the age
}

func main() {
	// Struct declaration without initialization
	var p1 People
	p1.Name = "John"
	p1.LastName = "Wick"
	p1.Age = 30
	p1.Status = true
	p1.documentID = "123.456.789-00" // Private field

	// Struct declaration with initialization
	p2 := People {
		Name: "Ethan",
		LastName: "Hunt",
		Age: 50,
		Status: true,
		documentID: "987.654.321-00",
	}

	// Struct declaration with initialization without field name
	var p3 People = People {
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

	// Print the structs using the show method
	p1.show()
	p2.show()
	p3.show()

	// Call the birthday method
	p1.birthday()
	p2.birthday()
	p3.birthday()

	// Print the structs using the show method
	p1.show()
	p2.show()
	p3.show()

	// Print the structs using the string method
	fmt.Println(p1.string())

	// Struct declaration with initialization
	cat1 := Category {
		Name: "Electronic",
	}

	// Struct declaration with initialization
	cat2 := Category {
		Name: "Sport",
		Father: &cat1,
	}

	// Check if the father is not nil
	if cat1.HasFather() {
		fmt.Println(cat1.Father.Name) // Print the father name
	} else {
		fmt.Println("No father") // Print no father
	}

	// Print the structs
	cat1.show()
	cat2.show()
	fmt.Println(cat1)
	fmt.Println(cat2)
}
