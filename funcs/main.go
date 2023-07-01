package main

import (
	"fmt"
	"strconv"
)

// func sum
func sum(x, y int) int {
	return x + y
}

// func subtract
func subtract(x, y int) int {
	return x - y
}

// func multiply
func multiply(x int, y int) int {
	return x * y
}

// func divide
func divide(x int, y int) int {
	return x / y
}

// func mod
func mod(x int, y int) int {
	return x % y
}

func convertAndSum(x int, y string) int {
	// convert string to int
	i, _ := strconv.Atoi(y) // _ is the error (ignored in thi)
	return x + i
}

func convertAndSum2(x int, y string) (result int, err error) {
	// convert string to int with error
	i, err := strconv.Atoi(y)

	if err != nil {
		return
	}

	result = x + i
	return
}

// Exported functions
// Sum exported, the first letter is uppercase
func Sum(x, y int) int {
	return x + y
}


func main() {
	// func sum
	fmt.Printf("sum: %d\r\n", sum(1, 2))
	fmt.Printf("subtract: %d\r\n", subtract(4, 2))
	fmt.Printf("multiply: %d\r\n", multiply(2, 2))
	fmt.Printf("divide: %d\r\n", divide(4, 2))
	fmt.Printf("mod: %d\r\n", mod(4, 2))
	fmt.Printf("convertAndSum: %d\r\n", convertAndSum(4, "2"))
	result, err := convertAndSum2(4, "2")
	fmt.Println("convertAndSum2: ", result, err)
}
