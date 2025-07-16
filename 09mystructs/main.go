package main

import 
(
	"fmt"
)

func main(){
	fmt.Println("Welcome to structs in golang")

	// struct is a user defined data type in golang
	// struct is a collection of fields
	// struct is a collection of data items
	// struct is a collection of methods
    
	jyotish := User{"Jyotishmoy", "jyotishmoy@go.dev", true, 21}
	fmt.Println(jyotish)
	fmt.Printf("Jyotish details are : %+v\n", jyotish)
	fmt.Println(jyotish.Name)
	fmt.Println(jyotish.Email)
	fmt.Println(jyotish.status)
	fmt.Println(jyotish.Age)

}

type User struct{
	Name string
	Email string
	status bool
	Age int
}