package main

import "fmt"

var address string
var phone string = "123-456-7890"
var name, lastname string = "John", "Doe"
var qtd int
var buy bool
var value float64
var words rune

func main () {
	fmt.Printf("addres: %s\r\n", address)
	fmt.Printf("phone: %s\r\n", phone)
	fmt.Printf("name: %s\r\n", name)
	fmt.Printf("lastname: %s\r\n", lastname)
	fmt.Printf("qtd: %d\r\n", qtd)
	fmt.Printf("buy: %t\r\n", buy)
	fmt.Printf("value: %f\r\n", value)
	fmt.Printf("words: %c\r\n", words)
}
