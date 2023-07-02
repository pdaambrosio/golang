package main

import (
	"fmt"
	// "os"
	// "log"
)

func isGreater(a, b int) string {
	if a > b {
		return "a is greater than b"
	} else if a < b {
	return "a is less than b"
	} else {
		return "a is equal to b"
	}
}

func main() {
	fmt.Println(isGreater(1, 2))
	fmt.Println(isGreater(2, 1))
	fmt.Println(isGreater(2, 2))
}
