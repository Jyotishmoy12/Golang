package main

import (
	"fmt"
)

func main (){
	fmt.Println("if else in golang!")

	// loginCount :=23
    // var result string 

	// if loginCount < 10 {
	// 	result = "regular user"
	// } else if loginCount >10 {
	// 	result = "premium user"
	// } else{
	// 	result = "admin user"
	// }

	// fmt.Println(result)


	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num :=3; num <10 {
		fmt.Println("Number is less than 10")
	} else{
		fmt.Println("Number is greater than 10")
	}

	
}