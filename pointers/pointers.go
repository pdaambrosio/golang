package main

import "fmt"

/*
Pointers are memory addresses of variables in memory (RAM)
Pointers are useful when you want to share a variable between two functions
Pointers are also useful when you want to change the value of a variable inside a function
*/

func main() {
	a := 10 // this is a variable
	b := &a // this is a pointer to a variable (memory address of a)

	var firstPointer *string // this pointer isn't initialized yet
	// firstPointer = "Mafalda" // this is a string literal, not a pointer to a string and it will throw an error

	var secondPointer *string = new(string) // this pointer is initialized with the new() function
	*secondPointer = "Mafalda" // this is a pointer to a string and it will work (desreferencing)

	thirdPointer := "Black Bozo" // this is a string literal
	pointToThirdPointer := &thirdPointer // this is a pointer to a string literal (address of thirdPointer)
	
	fmt.Println(a, b) // 10 0xc00009e230
	fmt.Println(a, *b) // 10 10 (desreferencing value)
	fmt.Println(firstPointer) // <nil> (zero value) - no memory address
	fmt.Println(secondPointer) // 0xc00009e230 (dynamic memory address)
	fmt.Println(*secondPointer) // Mafalda (desreferencing value)
	fmt.Println(thirdPointer) // Black Bozo (string literal)
	fmt.Println(pointToThirdPointer) // 0xc00009e240 (dynamic memory address)
	fmt.Println(pointToThirdPointer, *pointToThirdPointer) // Black Bozo (desreferencing value)
	thirdPointer = "Black Bozo 2" // this will change the value of thirdPointer and pointToThirdPointer, both has the same memory address
	fmt.Println(pointToThirdPointer, *pointToThirdPointer) // Black Bozo 2 (desreferencing value)
}
