package main

import (
	"fmt"
	"time"
)

// go routine is a lightweight thread of execution (a function) that run concurrently with other goroutines

// This is main function (entry point) of the program
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
		fmt.Println(s) // Print the string
	}
}

// This is main function (entry point) of the program
func main() {
	go say("world") // Create a new goroutine and run say("world") in it concurrently with the calling one
	say("hello") // Call say("hello") in the main goroutine
}
