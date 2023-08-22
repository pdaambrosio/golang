package main

import (
	"fmt"
	"time"
)

/*
 * The function main() creates a buffered channel of capacity 3.
 * Then, it starts a goroutine that sends 3 elements to the channel and closes it.
 * The main goroutine sleeps for 2 seconds and then starts receiving elements from the channel.
 * The main goroutine will receive all the elements and then it will receive a zero value.
 * The zero value is received because the channel is closed.
 * The main goroutine will break the for loop and print the message End.
*/

func main() {
	c := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sender: sending element %v...\n", i)
			c <- i
		}
		fmt.Println("Sender: close the channel...")
		close(c)
	}()

	time.Sleep(2 * time.Second)
	for {
		v, ok := <-c
		if !ok {
			fmt.Println("Receiver: channel is closed...")
			break
		}
		fmt.Printf("Receiver: received an element: %v\n", v)
	}
	fmt.Println("End.")
}
