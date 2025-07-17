package main

import "fmt"

func main(){ 
	// so basically defer is a stack of functions that will be executed in reverse order
	// last in first out LIFO 
	defer fmt.Println("This will print last!")
	defer fmt.Println("This will print second last!")
	defer fmt.Println("This will print first!")
	fmt.Println("Welcome to defer in golang!") 
	myDefer()

}

func myDefer (){
	for i :=0;i<5;i++{
		defer fmt.Println(i)
	}
}