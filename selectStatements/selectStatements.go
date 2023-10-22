package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "I'm channel 1"
	}()

	go func() {
		ch2 <- "I'm channel 2"
	}()

	time.Sleep(20 * time.Microsecond)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("No message received")
	}
}
