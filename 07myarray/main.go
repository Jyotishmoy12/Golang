package main

import "fmt"

func main() {

	fmt.Println("Welcome to arrays in golang")

	var fruitList [5] string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Cherry"
	fruitList[4] = "Date"
	
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegetableList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegetable list is: ", vegetableList)
	fmt.Printf("Vegetable list is: %T\n", vegetableList)

}