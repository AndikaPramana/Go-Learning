package chapter2

// this package is used to learn more in-depth about Go

import (
	"fmt"
)

func main() {

	// variable declaration
	// static
	var x int
	x = 20

	fmt.Println(x)

	// this will result an error ==> x = 20.5
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is type of %T\n\n", x)

	// Dynamic or can be said short assignment, without type
	y := 20.5
	fmt.Println(y)
	fmt.Printf("y is type of %T\n", y)

	y = 20
	fmt.Println(y)
	fmt.Printf("y is type of %T\n\n", y)

	// mixed variable declaration
	var a, b, c = 1, 10.5, "hello"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c, "\n")

	// for loop example
	// below example will result error, because variable type is not initialize
	// notice the difference of colon ":"
	/*
		for i = 0; i < 5; i++ {
			fmt.Printf("value of i: %d\n", i)
		}
	*/

	for i := 0; i < 5; i++ {
		fmt.Printf("value of i: %d\n", i)
	}

	numbers := [6]int{1, 2, 3, 5}

	// looping for array, similar to foreach
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

}
