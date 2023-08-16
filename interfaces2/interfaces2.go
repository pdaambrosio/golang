package main

import "fmt"

type Car struct {
	manufacturer string
	model string
	year int
}

type People struct {
	name string
	age int
	Car Car
}

func goToWork(people People) {
	fmt.Println(people.name, "is going to work with a", people.Car.manufacturer, people.Car.model)
}

func main() {
	var bmw Car = Car{
		manufacturer: "BMW",
		model: "X5",
		year: 2020,
	}

	var p1 People = People{
		name: "John",
		age: 20,
		Car: bmw,
	}

	fmt.Println(p1)
	goToWork(p1)
}