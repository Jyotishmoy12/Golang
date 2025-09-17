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


type Product struct{
	Name string
	Price int
	Company string
}

// func fuc(copyOfP Product){
// 	copyOfP.Name = "Mackbook pro" // it will not change the actual product struct it will be pass by reference
// }

// func newProduct(name string, price int, company string) Product{
// 	p:= Product{
// 		Name: name,
// 		Price: price,
// 		Company: company,
// 	}
// 	return p // this is returing a copy of the product struct
// }

func newProduct(name string, price int, company string) *Product{
	p:= Product{
		Name: name,
		Price: price,
		Company: company,
	}
	return &p // this is returing the actual product struct
}

func fun_pass_by_ref(p* Product){
	p.Name = "Mackbook pro"
}
// some example of copy usecases


// member functions use pass by reference
// member function of a struct

func (p *Product) display(){
   fmt.Println("Product details are: ", p.Name, p.Price, p.Company)
}

// call this function
// new_p := newProduct("Iphone", 10000, "Apple")
// new_p.display()
