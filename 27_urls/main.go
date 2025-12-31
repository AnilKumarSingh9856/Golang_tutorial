package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghb38kegs"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.RawPath)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Println(result.Query())


	qparams := result.Query()
	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is : ", val)
	}

	// constructing url from pieces
	partsOfUrl := &url.URL{  // keep in mind, always use '&' for reference not use used copy
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}