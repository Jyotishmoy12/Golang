package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax/ comma err syntax
	// it returns a value and an error like try catch in other languages
	// _ is a blank identifier which is used to ignore the error
	// _ is used when we don't care about the error
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	// its a string type 
	fmt.Printf("Type of this rating is %T", input)
}