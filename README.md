# golang

Trying to learn Go. This is a collection of notes and code snippets. (It's looks like a mix of C and Python :weary:)

## Installation

```bash
apt install golang
```

or

```bash
brew install golang
```

## Add to ~/.bash_profile

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

The first line tells Go to use the go directory in your home directory as the workspace. The second line tells your shell where to find the Go binaries.

It is recommended to create a directory for your Go projects, and set GOPATH to the path of this directory. For example, if you create a directory named go in your home directory, you should set GOPATH to /home/your-user-name/go.

## Hello World

```bash
mkdir hello
cd hello
touch hello.go
```

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

```bash
go run hello.go
```

## References

- [Go Website](https://golang.org/)
- [Go Documentation](https://golang.org/doc/)
- [Go Packages](https://golang.org/pkg/)
