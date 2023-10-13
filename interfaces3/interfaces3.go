package main

import (
	"bytes"
	"fmt"
	"strings"
)

// Interface is a type that specifies a method set called its interface
type printer interface {
	Print() string
}

// Struct user is a simple struct with two fields
type user struct {
	username string
	id       int
}

// Method Print of the user returns a string with the username and id
func (u user) Print() string {
	return fmt.Sprintf("%v (%v)", u.username, u.id)
}

// Struct menuItem is a simple struct with two fields
type menuItem struct {
	name  string
	price map[string]float64
}

// Method Print of the menuItem returns a string with the name and price
func (m menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%v\n", m.name))
	b.WriteString(strings.Repeat("-", 10) + "\n")
	for size, cost := range m.price {
		fmt.Fprintf(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}

// Function main is the entry point of the program
func main() {
	var p printer
	p = user{username: "John",id: 1}
	fmt.Println(p.Print())

	p = menuItem{
		name: "Pizza",
		price: map[string]float64{
			"small":  5.99,
			"medium": 7.99,
			"large":  9.99,
		},
	}
	fmt.Println(p.Print())

	// Type assertion
	u, ok := p.(user)
	fmt.Println(u, ok)
	
	m, ok := p.(menuItem)
	fmt.Println(m, ok)

	// Type switch
	switch v := p.(type) {
	case user:
		fmt.Println("User:", v.username)
	case menuItem:
		fmt.Println("Menu item:", v.name)
	default:
		fmt.Println("Unknown type")
	}
}
