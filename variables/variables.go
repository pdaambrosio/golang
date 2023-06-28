package main

import "fmt"

var address string
var phone string = "123-456-7890"
var name, lastname string = "John", "Doe"
var qtd int
var buy bool
var value float64
var words rune

var (
	information     string
	age             int
	Addres2, phone2 string //Address2 is public, phone2 is private
	password        string = "123456"
)


func main() {
	test := "test" //suggar sintax
	fmt.Printf("addres: %s\r\n", address)
	fmt.Printf("phone: %s\r\n", phone)
	fmt.Printf("name: %s\r\n", name)
	fmt.Printf("lastname: %s\r\n", lastname)
	fmt.Printf("qtd: %d\r\n", qtd)
	fmt.Printf("buy: %t\r\n", buy)
	fmt.Printf("value: %f\r\n", value)
	fmt.Printf("words: %c\r\n", words)
	fmt.Printf("information: %s\r\n", information)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("addres2: %s\r\n", Addres2)
	fmt.Printf("phone2: %s\r\n", phone2)
	fmt.Printf("test: %s\r\n", test)
}
