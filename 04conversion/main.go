package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/**
	Here in the below example we are accepting the user input and then type cast it
	we are using package bufio, os, strconv, strings for doing this
	bufio - this is buffer reader
	os - os pakcage is used for the using the stdin in os package
	steconv - this pkg is used for converting the converting the string
	strings - to trim the blank spaces use use TrimSpace method of package strings
	*/
	fmt.Println("Welcome to the Pizza rating app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Your input is ", input)
	fmt.Println("Thanks for your rating")

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
	fmt.Printf("Type of rating %T", input)

	rating, errr := strconv.ParseInt(strings.TrimSpace(input), 0, 36)
	fmt.Println("your rating is: ", rating)

	if errr != nil {
		fmt.Println("Error :- ", errr)
	} else {
		fmt.Println("Rating:- ", rating)
	}

}
