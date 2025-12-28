package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://leetcode.com/u/Anil_kumar_Singh/"

func main() {
	fmt.Println("LCO web request")
	 
	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	// datatypes, err := ioutil.ReadAll(response.Body)
	datatypes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(datatypes)
	fmt.Println(content)


}
