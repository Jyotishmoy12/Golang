package main

import "fmt"

// interfaces
type sellableProduct interface {
	buy()
	getDiscount() int
}


// struct
type Product struct {
	Name    string
	Price   int
	Company string
}

// constructor 
func newProduct(name string, price int, company string) *Product {
	p := Product{
		Name:    name,
		Price:   price,
		Company: company,
	}
	return &p // this is returing the actual product struct
}

func (p Product) buy() {
	fmt.Println("Buying the product", p.Name, "for Price", p.Price)
}
func (p * Product) getDiscount() int {
	discount := p.Price * 10 / 100
	fmt.Println("Discount for product", p.Name, "is", p.Price)
	return discount
}
func check_discount_and_buy(p sellableProduct) {
	discount := p.getDiscount()
	if(discount>30){
		fmt.Println("Discount is good buying the product")
		p.buy()
		return
	}else{
		fmt.Println("Discount is not good buying the product")
		return
	}
}
// member function use pass by reference
func (p *Product) display() {
	fmt.Println("Product details are: ", p.Name, p.Price, p.Company)
}



func main(){

	// call this function
	new_p := newProduct("Iphone", 10000, "Apple")
	new_p.display()

	// call this function
	check_discount_and_buy(new_p)
}