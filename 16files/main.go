package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	fmt.Println("Files in golang")

	// reading files in golang
	content :="This needs to go in a file - youtube.com"

	file,err := os.Create("./myfile.txt")

	if err !=nil{
		panic(err)
	} 

	length, err := io.WriteString(file, content)

	if err !=nil{
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string){
	// reading files in golang
	databyte, err := os.ReadFile(filename)

	if err !=nil{
		panic(err)
	} 

	fmt.Println("Text data inside the file is:\n ", string(databyte))


}