package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main(){
	fmt.Println("GET/POST requests in golang")
	//PerformGetRequest()
	//PerformPostRequest()
	PerformPostFormRequest()
}


func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)


	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Println("Response string is: ", responseString.String())

	// fmt.Println("Content: ", string(content))
}

func PerformPostRequest(){
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(

		`{
			"coursename": "Let's go with golang",
			"price": 0,
			"platform": "learncodeonline.in"
		}`,
		
	)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	// fake formdata

	data := url.Values{}
	data.Add("firstname", "jyotish")
	data.Add("lastname", "moy")
	data.Add("email", "jyotishmoy@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))
     
	
}