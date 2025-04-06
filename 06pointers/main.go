package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in GO")
	/**
	Why pointers ?
	Operations will be performed on actual values.
	When we create array or any other data structure, and pass them to multiple functions for operations and modifications then the copy of that array or ddata structure is passed, this cause irregularity and sometime its difficult to handled and denugged thats why points are exists in this case you pass the pointer (memory address) to the function or method and made those changes and modifications.

	& - get Memory address
	* - get the actual value
	*/

	var ptr *int
	fmt.Println("Value of pointer is :-", ptr)

	mynumber := 23

	var pointer = &mynumber
	fmt.Println("Value of pointer is :-", pointer)
	fmt.Println("Value of pointer is :-", *pointer)

	*pointer = *pointer * 2
	fmt.Println("New value is :-", mynumber)

}
