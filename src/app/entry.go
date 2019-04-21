package main

import (
	"fmt"
	// nested package
	"greet"
	"greet/de"
	"math"
	"math/rand"

	// package alias
	child "greet/greet"
)

func init() {
	// function to initialize global variable
	fmt.Println("THIS IS INIT FUNCTION\n\n")
}

// testing comment
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 4))

	// importing from a difference package
	fmt.Println(greet.Morning)
	// nested package
	fmt.Println(de.Morning)

	// because same namespace "app", we can just use the variable without import
	fmt.Println("version: " + version)

	/*variable initialization*/
	// not allowed to use variable before declaration, inside a package
	// but we can use it if it is defined in package scope
	/*
		var a int = b
		var b int = c
		var c int = 2

		if it is declared inside function, compile error
		declare it package scope and it will run just fine
	*/

	fmt.Println(a, b, c)

	// importing function between same name space
	fmt.Println(func_pack())

	// using package alias
	fmt.Println(child.Message)
}

var (
	a int = b
	b int = f()
	c int = 1
)

func f() int {
	return c + 1
}

// testing comment on another user commit
func add(x int, y int) int {
	return x + y
}
