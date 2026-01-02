package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// How to create json data
func EncodeJson() {

	lcoCourses := []Course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"full stack project"}},
		{"Golang Bootcamp", 289, "LearnCodeOnline.in", "bcd123", []string{"backend project"}},
		{"Rust Bootcamp", 399, "LearnCodeOnline.in", "efg123", nil},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

// How to consume json data
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 	{
			"coursename": "Golang Bootcamp",
			"Price": 289,
			"website": "LearnCodeOnline.in",
			"tags": ["backend project"]
        }
	`)

	var lcoCourses Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	}else{
		fmt.Println("JSON is invalid")
	}

	// some cases where you just want to add data to key value

	var myOnlineCourses map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineCourses)
	fmt.Printf("%#v\n", myOnlineCourses)

	for k, v := range myOnlineCourses{
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}

func main() {
	fmt.Println("Welcome to the JSON video.")

	// EncodeJson()
	DecodeJson()

}
