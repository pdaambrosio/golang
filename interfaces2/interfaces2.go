package main

import "fmt"

type People struct {
	name string
	age int
}

func main() {
	var p1 People = People{"zhangsan", 18}
	fmt.Println(p1)
}