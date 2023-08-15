package main

import "fmt"

// this is a simple structs
type People struct {
	Name string
	Age int
	Naturality string
	CPF string
}

type LegalPeople struct {
	People
	cnpj string // this is a private field
}

// this is a interface
type Document interface {
	Doc() string
}

// this is a interfaces that inherits from Document interface
func (p People) Doc() string {
	return p.CPF
}

func (lp LegalPeople) Doc() string {
	return lp.cnpj
}

// this is a function that receives a Document interface and assert if it is a LegalPeople struct
func show(d Document) {
	if d, ok := d.(LegalPeople); ok {
		fmt.Println(d.cnpj)
	}
}

// this is a function that receives a Document interface and assert if it is a People or LegalPeople struct
func show2(d Document) {
	switch d.(type) {
	case People:
		fmt.Println(d.(People).CPF)
	case LegalPeople:
		fmt.Println(d.(LegalPeople).cnpj)
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	people1 := People{"John", 30, "Brazilian", "123.456.789-00"} // creating a new People struct
	legalPeople1 := LegalPeople{people1, "12345678900000"} // creating a new LegalPeople struct

	fmt.Println(people1)
	fmt.Println(people1.Doc())
	fmt.Println(legalPeople1)
	fmt.Println(legalPeople1.Doc())

	show(legalPeople1)
	show2(people1)
	show2(legalPeople1)
}
