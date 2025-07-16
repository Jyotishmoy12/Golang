package main

import "fmt"

// capital letter means public variable and small letter means private
const LoginToken string = "jyotishmoy"

func main() {
	// var username string = "Jyotishmoy"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type: %T", username)

	// var smallVal uint8 = 255
	// fmt.Println(smallVal)
	// fmt.Printf("Variable is of type: %T", smallVal)

	// default value and aliases
	// var anotherVariable int
	// fmt.Println(anotherVariable)
	// fmt.Printf("Variable is of type: %T", anotherVariable)

	// implicit type
	// var website = "google.com"
	// fmt.Println(website)

	// no var style
	// outside the method or in global 
	// scope we are not allowed to declare varianble without var keyword
	numberOfUsers := 3000000
	fmt.Println(numberOfUsers)
   
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T", LoginToken)
}