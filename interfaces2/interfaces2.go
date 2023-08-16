package main

import "fmt"

type Vehicle interface {
	accelerate()
}

// Struct is a way to create a custom data type
type Car struct {
	manufacturer string
	model string
	year int
}

// This is a interface method implementation for Car struct
func (c Car) accelerate() {
	fmt.Println("The", c.manufacturer, c.model, "is accelerating")
}

// Struct is a way to create a custom data type
type Motorcycle struct {
	manufacturer string
	model string
	year int
}

// This is a interface method implementation for Motorcycle struct
func (m Motorcycle) accelerate() {
	fmt.Println("The", m.manufacturer, m.model, "is accelerating")
}

// Struct can be used as a field in another struct (composition)
type People struct {
	name string
	age int
	Car Car
	Motorcycle Motorcycle
	Venicle Vehicle
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

	var yamaha Motorcycle = Motorcycle{
		manufacturer: "Yamaha",
		model: "XJ6",
		year: 2019,
	}

	var john People = People{
		name: "John",
		age: 20,
		Car: bmw,
	}

	var stalone People = People{
		name: "Stalone",
		age: 50,
		Motorcycle: yamaha,
	}

	var link People = People{
		name: "Link",
		age: 30,
		Venicle: yamaha,
	}

	var zelda People = People{
		name: "Zelda",
		age: 30,
		Venicle: bmw,
	}

	fmt.Println(john) // print the struct
	goToWork(john)  // call the simple function
	john.goToWorkAgain() // call the method
	goToHome(&john) // call the function that receives a pointer to a struct
	fmt.Println(stalone) // print the struct again
	stalone.goToWorkAgain() // call the method again but with another struct without car (missing fields)
	fmt.Println(link) // print the struct again
	link.Venicle.accelerate() // call the method again but with another struct
	fmt.Println(zelda) // print the struct again
	zelda.Venicle.accelerate() // call the method again but with another struct
}