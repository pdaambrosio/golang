package main

import "fmt"
import "strconv"

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


// Constants
const (
	PI = 3.14
	Language = "Go"
)

// Constants with iota

const (
	_ = iota // 0
	KB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	MB = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
	GB = 1 << (10 * iota) // 1 << (10 * 3) = 1073741824
	TB = 1 << (10 * iota) // 1 << (10 * 4) = 1099511627776
	PB = 1 << (10 * iota) // 1 << (10 * 5) = 1125899906842624
	EB = 1 << (10 * iota) // 1 << (10 * 6) = 1152921504606846976
	ZB = 1 << (10 * iota) // 1 << (10 * 7) = 1180591620717411303424
	YB = 1 << (10 * iota) // 1 << (10 * 8) = 1208925819614629174706176
)

// iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
// Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. Its value is the index of the respective ConstSpec in that constant declaration, starting at zero. It can be used to construct a set of related constants:

// const (
// 	c0 = iota  // c0 == 0
// 	c1 = iota  // c1 == 1
// 	c2 = iota  // c2 == 2
// )

// const (
// 	a = 1 << iota  // a == 1  (iota == 0)
// 	b = 1 << iota  // b == 2  (iota == 1)
// 	c = 3          // c == 3  (iota == 2, unused)
// 	d = 1 << iota  // d == 8  (iota == 3)
// )

// convertion

// int to float
var i int = 42
var f float64 = float64(i)

// float to int
var f2 float64 = 42.123
var i2 int = int(f2)

// string to int, need to import strconv
var strC1 string = "42"
var intC1, err = strconv.ParseInt(strC1, 10, 64) // 10 is the base, 64 is the bit size

// string to float, need to import strconv
var strC2 string = "42.123"
var floatC2, err2 = strconv.ParseFloat(strC2, 64) // 64 is the bit size

// int to string, need to import strconv
var strC3 string = strconv.Itoa(42)

// float to string, need to import strconv
var strC4 string = strconv.FormatFloat(42.123, 'f', 6, 64) // 'f' is the format, 6 is the precision, 64 is the bit size

// boolean to string, need to import strconv
var strC5 string = strconv.FormatBool(true)

// string to boolean, need to import strconv
var boolC6, err6 = strconv.ParseBool("true")

// string to byte
var byteC7 = []byte("test")

// byte to string
var strC8 = string([]byte{116, 101, 115, 116})

// rune to string
var strC9 = string([]rune{116, 101, 115, 116})

// string to rune
var runeC10 = []rune("test")

var string1 = "this is a string" // This is a interpreted string (string interpolation) and can be used with \n, \t, etc
var string2 = `this is a string` // This is a raw string and if you use \n, \t, etc, this escape characters will be interpreted as commum text

// type conversions

var d float64 = 3.1415
// var e int = d // error, go don't allow implicit type conversions
var e int = int(d) // ok, this is a explicit type conversion

func main() {
	test := "test" // suggar sintax
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
	fmt.Printf("Int convert: %d\r\n", intC1)
	fmt.Printf("Float convert: %d\r\n", e)
}
