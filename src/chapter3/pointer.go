package main

// this package is used to learn more about Go pointer

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	// printing address of Go Memory
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a)

	var b int = 20
	var ip *int

	ip = &b

	fmt.Printf("Address of b variable: %x\n", &b)
	fmt.Printf("Value of b variable: %i\n", b)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	// when initialize pointer, value is 0, but address is not 0
	var c *int
	fmt.Printf("Address of c variable: %x\n", &c)
	fmt.Printf("Value of c variable: %d\n", c)

	// checking for nil pointer
	if c != nil {
		fmt.Println("Pointer c is not nil")
	} else {
		fmt.Println("Pointer c is nil")
	}

	fmt.Println("\n")
	//example of struct pointer

	var Book1 Books

	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	printBook(&Book1)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
