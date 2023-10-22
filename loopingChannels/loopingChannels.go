package main

import "fmt"

func main() {
	ch := make(chan int) // create a channel of type int (bidirectional channel)

	// Function literal (anonymous function) that sends 10 values to the channel and then closes the channel
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // close the channel after the loop is done to avoid deadlock (fatal error: all goroutines are asleep - deadlock!)
	}()

	// For loop that receives the values from the channel and prints them out until the channel is closed (range) and then exits the loop (break)
	for n := range ch {
		fmt.Println(n)
	}
}
