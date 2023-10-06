package main

import "fmt"

/*
 * panicAndRecover.go
 * Demonstrates how to use panic and recover
 * A panic is a situation when a program cannot continue to run because of an unrecoverable error.
 * A panic can be caused by runtime errors such as array out of bounds access, trying to divide by zero, or trying to allocate too much memory.
 * A panic can also be caused by programmer errors such as passing an invalid argument to a function or failing to check for errors.
 * When a panic occurs, the program immediately stops execution and prints the panic message.
 * The panic message is usually the error value that caused the panic.
 * After printing the panic message, the program terminates.
 * The panic message is printed to standard error.
 * The recover function can be used to recover from a panic.
*/

// The function "divide" divides two integers and handles any potential division by zero errors.
func divide(dividend, divisor int) int {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()

	return dividend / divisor
}

// The main function divides two numbers and prints the result, handling the case of dividing by zero.
func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))
}
