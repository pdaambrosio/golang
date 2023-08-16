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

func main() {
	var bmw Car = Car{
		manufacturer: "bmw",
		model: "x5",
		year: 2020,
	}

	var p1 People = People{
		name: "John",
		age: 20,
		Car: bmw,
	}
	
	fmt.Println(p1)
}