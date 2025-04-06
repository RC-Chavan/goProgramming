package main

import "fmt"

const Abc string = "ABCABC"

// always first letter should be capital if you are declaring any variable in global or outside main function
// Public variable - first name should be capital
// If you are declaring the varible in public you are not allowed to use no var type i.e. abc := "123123"
var jwtToken int = 4532323

func main() {
	var username string = "Rohit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isvarified bool = true
	fmt.Println(isvarified)
	fmt.Printf("Variable is of type: %T \n", isvarified)

	var smallvalue int = 10
	fmt.Println(smallvalue)
	fmt.Printf("Variable is of type: %T \n", smallvalue)

	var smallfloat float32 = 2555.3453454545453453
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	var bigfloat float64 = 2555.3453454545453453
	fmt.Println(bigfloat)
	fmt.Printf("Variable is of type: %T \n", bigfloat)

	var anothervariable int // default value is zero
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of type: %T \n", anothervariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type of declaring the variable

	var website = "rohitchavan.in"
	fmt.Println(website)

	// no var type of declaring the variable

	id := 4532323
	fmt.Println(id)

	fmt.Println(Abc)
	fmt.Printf("This is of type %T\t", Abc)
	fmt.Println(jwtToken)
	fmt.Printf("This is of type %T\t", jwtToken)
}
