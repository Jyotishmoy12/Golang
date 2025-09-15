package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is: ", ptr)
	fmt.Println("Value of actual pointer is: ", *ptr)

	*ptr = *ptr + 2
	// fmt.Println("Value of actual pointer is: ", *ptr)
	fmt.Println("Value of actual pointer is: ", myNumber)

	demo_pointers()
}

func demo_pointers() {

	i := 120
	var ptr *int = &i // this is a pointer to int type

	fmt.Println("Value of i is: ", i)
	fmt.Println("Value of ptr is: ", ptr)
	fmt.Println("Value of ptr is: ", *ptr)

}
