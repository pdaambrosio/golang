package main

import "fmt"

// String can't to be NULL
var address string
var phone string = "123-456-7890"
var name, lastname string = "John", "Doe"

// Boolean
var buy bool

// Integers
var qtd int // 32 bits or 64 bits
var qtd2 int8 // 8 bits -128 to 127
var qtd3 int16 // 16 bits -32768 to 32767
var qtd4 int32 // 32 bits -2147483648 to 2147483647
var qtd5 int64 // 64 bits -9223372036854775808 to 9223372036854775807
var qtd6 uint // 32 bits or 64 bits (unsigned)
var qtd7 uint8 // 8 bits 0 to 255 (unsigned)
var qtd8 uint16 // 16 bits 0 to 65535 (unsigned)
var qtd9 uint32 // 32 bits 0 to 4294967295 (unsigned)
var qtd10 uint64 // 64 bits 0 to 18446744073709551615 (unsigned)
var qtd11 uintptr // big enough to hold the bit pattern of any pointer

// Rune
var words rune // alias for int32, but the value expected is a unicode character

// Byte
var value byte // alias for uint8, but the value expected is a byte

// Float
var value1 float32 // 32 bits (IEEE-754)
var value2 float64 // 64 bits (IEEE-754)

// Complex
var value3 complex64 // 32 bits for imaginary part and an alias for float32, example: 12.8i
var value4 complex128 // 64 bits for imaginary part and an alias for float64, example: 12.8i

// By not declaring the data type, the compiler will automagically define a type for that variable

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
	fmt.Printf("value: %f\r\n", value1)
	fmt.Printf("words: %c\r\n", words)
	fmt.Printf("information: %s\r\n", information)
	fmt.Printf("age: %d\r\n", age)
	fmt.Printf("addres2: %s\r\n", Addres2)
	fmt.Printf("phone2: %s\r\n", phone2)
	fmt.Printf("test: %s\r\n", test)
}
