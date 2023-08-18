package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 10; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func letters() {
	for i := 'a'; i < 'z' + 10; i++ {
		time.Sleep(230 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers()
	go letters()
	time.Sleep(3 * time.Second)
	fmt.Println("main terminated")
}
