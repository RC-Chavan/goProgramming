package main

import "fmt"

func main() {
	/**
	If you are using array in Go, you must be mentioned the value while initialization of an array.
	*/
	fmt.Println("Welcome to Array in Go")

	var fruitList [4]string
	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	// fruitList[2] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println("Fruitlist is: ", fruitList)
	fmt.Println("Fruitlist is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("VegList is: ", len(vegList))
}
