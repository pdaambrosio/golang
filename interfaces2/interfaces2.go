package main

import "fmt"

// Struct is a way to create a custom data type
type Car struct {
	manufacturer string
	model string
	year int
}

// Struct can be used as a field in another struct (composition)
type People struct {
	name string
	age int
	Car Car
}

// This is a simple function that receives a struct as parameter
func goToWork(people People) {
	fmt.Println(people.name, "is going to work with a", people.Car.manufacturer, people.Car.model)
}

// This is a method that receives a struct as receiver
func (p People) goToWorkAgain() {
	fmt.Println(p.name, "is going to work again with a", p.Car.manufacturer, p.Car.model)
}

// This function receives a pointer to a struct as parameter, so it can change the struct
func goToHome(people *People) {
	people.Car.manufacturer = "Fiat"
	people.Car.model = "Uno"	
	people.name = "Snow"
	fmt.Println(people.name, "is going to home with a", people.Car.manufacturer, people.Car.model)
}

// This is main function (entry point) of the program
func main() {
	var bmw Car = Car{
		manufacturer: "BMW",
		model: "X5",
		year: 2020,
	}

	var john People = People{
		name: "John",
		age: 20,
		Car: bmw,
	}

	fmt.Println(john)
	goToWork(john)
	john.goToWorkAgain()
	goToHome(&john)
}