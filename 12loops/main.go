package main

import "fmt"

func main(){
	fmt.Println("Welcome to loops in golang")
    
	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "saturday"}
	
	fmt.Println(days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("Index is: %v and value is: %v\n", index, day)
	// }
	
	roughValue :=1

	for roughValue<10{

		if roughValue==5{
			roughValue++
			continue
		}
		fmt.Println(roughValue)
		roughValue++
	}
}