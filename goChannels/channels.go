package main

import (
	"fmt"
	"time"
)

func say(s string, done chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond) // Sleep for 100 milliseconds
		fmt.Println(s) // Print the string
	}

	done <- "done"
}

func main() {
	done := make(chan string) // Create a channel
	go say("world", done) // Create a new goroutine and run say("world") in it concurrently with the calling one
	// say("hello", done) // Call say("hello") in the main goroutine

	fmt.Println(<-done) // <-done blocks until it receives a value from done channel and then prints it
}