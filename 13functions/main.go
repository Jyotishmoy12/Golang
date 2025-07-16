package main

import "fmt"

func greeter() {
	fmt.Println("Hello, World!")
}

// func adder(valOne int, valTwo int, val3 string) (int, string) {
// 	return valOne + valTwo, val3
// }

func proAdder(values ... int) int{
	total :=0
	for _, val :=range values{
		total+=val
	}
	return total
}

func main() {
	fmt.Println("Welcome to functions in golang!")
	greeter()

	// result, message := adder(3, 5, "nice")

	// fmt.Println("Result is:", result, "Message is:", message)

	sum := proAdder(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50)
	fmt.Println("Sum is:", sum)
}
