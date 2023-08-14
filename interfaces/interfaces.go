package main

import "fmt"

// this is a simple struct
type Author struct {
	Name string
	Age int
	Naturality string
}

type Book struct {
	Author // inheritance from Author struct (composition)
	Title string
	Sinopse string
	Readed bool
}

// this is a method of Book struct that inherits from Readable interface (Readable interface)
func (b Book) Read() bool {
	return b.Readed
}

// this is a method of Author struct that inherits from Writable interface (Writable interface)
func (a Author) Write() string {
	return "Author: " + a.Name
}

// this is a method of Book struct that inherits from Readable and Writable interfaces (ReadWritable interface)
func (b Book) ReadWritable() string {
	return "Readed: " + fmt.Sprint(b.Read()) + " - " + b.Write()
}

// this is a interface (it's not necessary to use the word interface) that has a method Read()
type Readable interface {
	Read() bool
}

// this is a interface (it's not necessary to use the word interface) that has a method Write()
type Writable interface {
	Write() string
}

// this is a interface (it's not necessary to use the word interface) that inherits from Readable and Writable interfaces (ReadWritable interface)
type ReadWritable interface {
	Readable
	Writable
}

// this is the main function that will be executed when the program starts
func main() {
	author1 := Author{"John", 30, "Brazilian"} // creating a new Author struct
	book1 := Book{author1, "The book", "This is a book", false} // creating a new Book struct

	fmt.Println(book1)
	fmt.Println(book1.Read())
	fmt.Println(book1.Write())
	fmt.Println(book1.ReadWritable())
}
