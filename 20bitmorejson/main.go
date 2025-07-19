package main

import (
	"encoding/json"
	"fmt"
)


type course struct{
	Name string `json:"coursename"`
	Price int 
	Platform string `json:"website"`
	Password string `json:"-"`  // i dont want this field to be reflected in json 
	Tags []string `json:"tags,omitempty"` // omitempty --> if the value is empty, dont show it
}
func main(){
	fmt.Println("Lets learn how to create json in golang")
    // EncodeJson()
	DecodeJson()

}

func EncodeJson(){
	lcoCourses := []course{
		{
			"ReactJs Bootcamp", 
			299, 
			"learncodeonline.in", 
			"abc123", 
			[]string{"web-dev", "js"},
		},
		{
			"Mern Bootcamp", 
			199, 
			"learncodeonline.in", 
			"bcd123", 
			[]string{"full-stack", "js"},
		},
		{
			"Angular Bootcamp", 
			299, 
			"learncodeonline.in", 
			"hit123", 
			nil,
		},
	}
	// package this data into json data
	finalJson, err := json.MarshalIndent(lcoCourses,"", "\t")

	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson(){
	// decode json data into go struct
	jsonDataFromWeb := []byte(`
	   {
                "coursename": "ReactJs Bootcamp",
                "Price": 299,
                "website": "learncodeonline.in",
                "tags": ["web-dev","js"]
        }

	`)

	// how to verify if its actually a json data
	var lcoCourses course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("its a valid json")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)  // %#v--> for interfaces
	}else{
		fmt.Println("its not a valid json")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k,v :=range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v,v)
	}

}