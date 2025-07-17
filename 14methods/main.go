package main

import 
(
	"fmt"
)


func (u User) GetStatus(){
    fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail(){
	u.Email = "jyotishmoydeka@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
func main(){
	fmt.Println("Welcome to structs in golang")

	// struct is a user defined data type in golang
	// struct is a collection of fields
	// struct is a collection of data items
	// struct is a collection of methods
    
	jyotish := User{"Jyotishmoy", "jyotishmoy@go.dev", true, 21}
	// fmt.Println(jyotish)
	// fmt.Printf("Jyotish details are : %+v\n", jyotish)
	// fmt.Println(jyotish.Name)
	// fmt.Println(jyotish.Email)
	// fmt.Println(jyotish.Status)
	// fmt.Println(jyotish.Age)

	jyotish.GetStatus()
	jyotish.NewMail()

}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

