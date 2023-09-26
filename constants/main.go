package main

import "fmt"

/*
  Constants are declared like variables, but with the const keyword.
  Constants can be character, string, boolean, or numeric values.
  Constants cannot be declared using the := syntax.
  Constants can't be reassigned.
  An untyped constant takes the type needed by its context.
*/

// The line `const int = 42` is declaring a constant named `int` with a value of 42.
const num1 int = 42

// List of constants
const (
	num2 int = 42
	num3 float32 = 42.78
	str1 string = "Hello"
	str2 string = "World"
)


// In this case 'b' will have the same value as 'a' (2 * 5) because it's not initialized.
const (
	a = 2 * 5
	b // 2 * 5
)

/*
  iota is a predeclared identifier representing the untyped integer ordinal number
  of the current const specification in a (usually parenthesized) const declaration.
  iota starts at 0 and increments by one for each item in the list.
  iota resets to 0 whenever the word const appears in the source code,
  and increments after each const specification.
*/

// iota increments by 1 in each line of the block
const (
	c = iota // c == 0
	d = iota // d == 1
)

// iota is reset to 0
const e = iota // e == 0

// iota value depends on the line where it's declared
const (
	f = iota // f == 0
	g = 2.1 // g == 2.0
	h = iota // h == 2
	i // i == 3
)

// Example of iota usage
// KB, MB, GB, TB, PB, EB, ZB, YB
type ByteSize float64

// The code block `const (
//     _           = iota // ignore first value by assigning to blank identifier
//     KB ByteSize = 1 << (10 * iota)
//     MB
//     GB
//     TB
//     PB
//     EB
//     ZB
//     YB
// )` is defining a set of constants representing different byte sizes.
const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

// The `func (b ByteSize) String() string` is a method defined on the `ByteSize` type. It allows you to
// define how the `ByteSize` type should be formatted when converted to a string.
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

// The main function prints out various variables and constants, including numbers, strings, and byte
// sizes.
func main() {
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(str1 + " " + str2)
	fmt.Println(a, b)
	fmt.Println(c, d, e)
	fmt.Println(f, g, h, i)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB) // 1e+03 1e+06 1e+09 1e+12 1e+15 1e+18 1e+21 1e+24
	var size ByteSize = 1e13 // 9.09TB (1e13 bytes) it's the same as 9.09 * 1024 * 1024 * 1024 * 1024
	fmt.Println(size)
}
