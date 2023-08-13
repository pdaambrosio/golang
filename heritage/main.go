package main

import "fmt"

type People struct {
	Name string
	Age int
	State bool
}

type NaturalPeople struct {
	People 	// inheritance from People struct (composition)
	Lastname string
	cpf string // private
}

type LegalPeople struct {
	People // inheritance from People struct (composition)
	CompanyName string
	cnpj string // private
}

func (p NaturalPeople) String() string {
	return "Name: " + p.Name + " " + p.Lastname + "\nAge: " + fmt.Sprintf("%d", p.Age) + "\nState: " + fmt.Sprintf("%t", p.State) + "\nCPF: " + p.cpf
}

func (l LegalPeople) String() string {
	// Sprintf is used to convert the int and bool values to string
	return fmt.Sprintf("Name: %s\nAge: %d\nState: %t\nCNPJ: %s", l.Name, l.Age, l.State, l.cnpj)
}

// alternative way to print the struct
func (n NaturalPeople) String2() string {
	// n.People.Name is used to access the People struct fields (composition) but it's not necessary to use it
	return fmt.Sprintf("Name: %s\nAge: %d\nState: %t\nCPF: %s\nLastname: %s", n.People.Name, n.People.Age, n.People.State, n.cpf, n.Lastname)
}

func main() {
	np1 := NaturalPeople{People{"John", 30, true}, "Connors", "123.456.789-00"}
	lp1 := LegalPeople{People{"John", 30, true}, "Connors", "123.456.789-00"}

	fmt.Println(np1)
	fmt.Println(lp1)
}
