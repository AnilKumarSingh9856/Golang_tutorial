package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)


func main() {
	fmt.Println("Welcome back to handle Get request")

	// PerformGetRequest()
	// PerformPostRequest()
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
	fmt.Println("Content length is: ", response.ContentLength)

	// content, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// // // 1st way
	// // fmt.Println(content)
	// // fmt.Println(string(content))


	// 2nd way
	var responseString strings.Builder
	content2, err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}

	byteCount, _ := responseString.Write(content2)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(responseString.String())
}


func PerformPostRequest(){
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
		"coursename": "Let's go with golang",
		"price": 0,
		"plateform": "learnCodeOnline"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err !=nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}

	data.Add("firstname", "Anil")
	data.Add("middlename", "Kumar")
	data.Add("lastname", "Singh")

	response, err := http.PostForm(myurl, data)
	if err !=nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
