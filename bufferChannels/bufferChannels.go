package main

import (
	"fmt"
	"time"
)

// with out buffer channels, the main function will wait for the function numbers to send a value to the channel ch
// function numbers will receive a channel only for sending values to it (chan <- int) and will send 10 values to it (0 to 9)
// delay 90 milliseconds between each value sent
func numbers(n chan <- int) {
	for i := 0; i < 10; i++ {
		n <- i // will write the value i to the channel n
		fmt.Printf("Writen in Channel: %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Printf("End of writing\n")
	close(n) // will close the channel n after the loop is done (10 values sent)
}

// main function will create a channel of type int and will call the function numbers in a goroutine
func main() {
	ch := make(chan int) // create a channel of type int (without buffer)
	go numbers(ch) // will call the function numbers in a goroutine

	for r := range(ch) { // will loop through the channel ch
		fmt.Printf("Readed from the Channel: %d\n", r)
		time.Sleep(time.Millisecond * 130)
	}

	fmt.Print("End of CH\n")
	fmt.Printf("\n\n\n")

	// buffer channels will not wait for the main function to read the values sent to the channel ch2 (will not wait for the loop)
	ch2 := make(chan int, 3) // create a channel of type int with buffer of 3 values (3 is length of the buffer)
	go numbers(ch2) // will call the function numbers in a goroutine

	for r := range(ch2) { // will loop through the channel ch2
		fmt.Printf("Readed from the Channel: %d\n", r)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Print("End of CH2\n")
	fmt.Printf("End of Program\n")
}
