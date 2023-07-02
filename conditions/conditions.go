package main

import (
	"fmt"
	"os"
	"log"
)

func isGreater(a, b int) string {
	// If-else statement
	if a > b {
		return "a is greater than b"
	} else if a < b {
	return "a is less than b"
	} else {
		return "a is equal to b"
	}
}


func readFile(filename string) string {
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Read file, 100 bytes at a time
	data := make([]byte, 100)
	// Read up to 100 bytes
	if _, err := file.Read(data); err != nil {
		log.Fatal(err)
	}

	// Convert bytes to string
	return string(data)
}

func writeFile(filename string, data string) {
	// Open file for writing
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// Write string to file
	file.WriteString(data)
}

func isGreaterSwitch(a, b int) string {
	// Switch statement
	switch {
	case a > b:
		return "a is greater than b"
	case a < b:
		return "a is less than b"
	default:
		return "a is equal to b"
	}
}


func main() {
	fmt.Println(isGreater(1, 2))
	fmt.Println(isGreater(2, 1))
	fmt.Println(isGreater(2, 2))
	fmt.Println(readFile("hello.txt"))
	writeFile("hello2.txt", "Hello, World2!")
	fmt.Println(isGreaterSwitch(1, 2))
}
