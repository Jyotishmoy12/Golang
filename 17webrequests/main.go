package main

import (
	"fmt"
	"io"
	"net/http"
)


const url ="https://youtube.com"
func main(){
	fmt.Println("Welcome to web requests in golang!")
     
	response,err :=http.Get(url)

	if err != nil{
		panic(err)
	}
	
	defer response.Body.Close() // close the response body to prevent memory leak its our reponsibility

	databytes, err := io.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}