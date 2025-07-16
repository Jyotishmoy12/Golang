package main

import "fmt"

func main(){
	fmt.Println("Welcome to maps in golang!")

	languages :=make(map[string]string)

	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"
	languages["go"] = "Golang"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("List of all languages: ", len(languages))
	fmt.Printf("Map is of type: %T\n", languages)
	fmt.Println("The value for ruby is: ", languages["rb"])
	delete(languages, "rb")
	fmt.Println("The value for ruby is: ", languages["rb"])
	fmt.Println("List of all languages: ", languages)

	// looping over maps in golang
	for key, value :=range languages{
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}