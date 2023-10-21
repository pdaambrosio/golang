package main

import (
	"fmt"
	"sync"
)

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.
*/

func main() {
	var wg sync.WaitGroup
	ch := make(chan string) // Create a channel of strings.

	wg.Add(1) // Add 1 to the WaitGroup counter.
	go func() {
		ch <- "Do some async thing" // Send a string to the channel.
	}()

	go func() {
		fmt.Println(<-ch) // Receive the string from the channel.
		wg.Done()         // Subtract 1 from the WaitGroup counter.
	}()
}
