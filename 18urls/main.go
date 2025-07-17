package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://youtube.com:3000/learn?jyotishmoy=reactjs&paymentid=hfehfioheoi"
func main(){
	fmt.Println("Welcome to url handling in golang!")
	fmt.Println(myurl)

	// parsing the url

	  result, _ := url.Parse(myurl)
	//   fmt.Println(result.Scheme)
	//   fmt.Println(result.Host)
	//   fmt.Println(result.Path)
	//   fmt.Println(result.Port())
	  fmt.Println(result.RawQuery)

	  // query params
	  queryParams := result.Query()
	  fmt.Printf("The type of query params are: %T\n", queryParams)

	  fmt.Println(queryParams["jyotishmoy"])

	  
	  

}