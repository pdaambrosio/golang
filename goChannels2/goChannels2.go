package main

import (
	"fmt"
	"time"
)

func numbers(done chan <- bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	done <- true
}

func letters(done chan <- bool) {
	for i := 'a'; i < 'l'; i++ {
		time.Sleep(230 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	done <- true
}

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go numbers(done1)
	go letters(done2)

	<-done1
	<-done2

	fmt.Println("main terminated")
}
