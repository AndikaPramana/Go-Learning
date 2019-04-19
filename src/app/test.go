package main

import (
	"fmt"
	"math"
	"math/rand"
)

// testing comment
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 4))
}

// testing comment on another user commit
func add(x int, y int) int {
	return x + y
}
