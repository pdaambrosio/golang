package main

import (
	"fmt"
	"sync"
)

/*
WaitGroups are used to wait for the program to finish goroutines.
*/

// Function wgFunc1 is a example without WaitGroup in goroutine.
func wgFunc1() {
	go func() {
		fmt.Println("Do some async thing") // Don't will be printed because the program will finish before.
	}()
}

// Function wgFunc2 is a example of WaitGroup with goroutine.
func wgFunc2() {
	var wg sync.WaitGroup
	wg.Add(1) // Add 1 to the WaitGroup counter.
	go func() {
		fmt.Println("Do some async thing")
		wg.Done() // Subtract 1 from the WaitGroup counter.
	}()
	wg.Wait() // Wait until the WaitGroup counter is 0.
}

func main() {
	wgFunc1()
	wgFunc2()
}
