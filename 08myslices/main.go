package main

import (
	"fmt"
	// "sort"
)

func main(){
	fmt.Println("Welcome to slices in golang")

	// var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// fruitList = append(fruitList, "Mango", "Banana")
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)	
    
	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 945
	// highScores[2] = 465
	// highScores[3] = 867
	// fmt.Println(highScores)
	// highScores = append(highScores, 555, 666, 321)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// how to remove value from slice based on its index

	var courses =[]string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
    
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
    


}