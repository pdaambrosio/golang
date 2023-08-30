package main

import (
	"fmt"
)


// for loop
func multiplicationTable() {
	// for init; condition; post {}
	for i := 1; i <= 10; i++ {
		// nested for loop
		for j := 1; j <= 10; j++ {
			// %d is a placeholder for a number
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		// print a new line
		fmt.Println()
	}
}


// for loop with slice
func loopSlice(dataArr []string) {
	// for init; condition; post {}
	for i := 0; i < len(dataArr); i++ {
		fmt.Println(dataArr[i])
	}
}


// for loop with slice and index
func loopSliceIndex(dataArr []string) {
	// for index, value := range dataArr {}
	for index, value := range dataArr {
		fmt.Println(index, value)
	}
}


// for loop with slice without index
func loopSliceWithoutIndex(dataArr []string) {
	// for _, value := range dataArr {}
	for _, value := range dataArr {
		fmt.Println(value)
	}
}

func infiniteLoop() {
	// for {}, go don't have while loop
	var i int
	for {
		fmt.Println("infinite loop")
		fmt.Println(i)
		i++
	}
}


func main() {
	// data array of string
	data := []string{"a", "b", "c"}
	// call function
	multiplicationTable()
	loopSlice(data)
	loopSliceIndex(data)
	loopSliceWithoutIndex(data)
	// infiniteLoop()
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for key, value := range wellKnownPorts {
		fmt.Println(key, value)
	}
}
