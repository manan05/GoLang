package main

import "fmt"

// firstLetter capital means it is a public variable
const LoginToken string = "sjaprqbvssjd"

func main() {
	var username string = "manan"
	fmt.Println("Name:", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println("Name:", isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println("Name:", smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.54588889666666666666
	fmt.Println("Name:", smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	var anotherString string
	fmt.Println("Default: ", anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// another way of declaring the variables
	// implicit way
	var website = "https://github.com/manan05"
	fmt.Println(website)

	// no var style
	numberOfUser := 50
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
