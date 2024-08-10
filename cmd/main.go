package main

import (
	"fmt"
	"go-by-test/internal/hello"
	"go-by-test/internal/integer"
	"go-by-test/internal/iterator"
)

func main() {

	fmt.Println(hello.Hello("edvaldo"))
	fmt.Println(integer.Add(2, 2))
	fmt.Println(iterator.Repeat("a"))

}
