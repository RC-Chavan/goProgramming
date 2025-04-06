package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to go Program to accept user input"
	fmt.Println(welcome)

	/**
	creating the reader from bufio package using NewReader function and using the os's stdin for accepting the values
	**/

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your Name")

	/**
	reader.ReadString gives two parameters
	1. The Value
	2. Error
	which is also known as comma ok, err ok
	I fyou dont want to cosider error you can use _
	**/

	input, _ := reader.ReadString('\n')

	fmt.Println("Your name is:", input)
	fmt.Printf("Type of input is %T", input)

}
