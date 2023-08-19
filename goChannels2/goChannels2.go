package main

import (
	"fmt"
	"time"
)

// The function numbers will send a value to the channel done
func numbers(done chan <- bool) { // send only channel (send to channel)
	for i := 0; i < 10; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	done <- true // send a value to the channel
}

// The function letters will send a value to the channel done
func letters(done chan <- bool) { // send only channel (send to channel)
	for i := 'a'; i < 'l'; i++ {
		time.Sleep(230 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	done <- true // send a value to the channel
}

func main() {
	done1 := make(chan bool) // create a channel of type bool (send only channel)
	done2 := make(chan bool) // create a channel of type bool (send only channel)

	go numbers(done1) // call the function numbers(done1)
	go letters(done2) // call the function letters(done2)

	<-done1 // receive a value from the channel done1
	<-done2 // receive a value from the channel done2

	fmt.Println("main terminated")
}
