package main

import "fmt"


// iterating over data structures
func main() {
	var a = []int{1,2,3,4}

	for i, v := range a {  // it give index and value
		fmt.Println(i, v)
	}

	for _, v := range a {  // if we want only value
		fmt.Println(v)
	}
	for v := range a {  // it give by default only index
		fmt.Println(v)
	}

	m := map[string]string{"fname":"Anil Kumar", "lname":"Singh"}

	for k, v := range m { // it give key and value
		fmt.Println(k, v)
	}

	for _, v := range m {  // if we want only value
		fmt.Println(v)
	}
	for v := range m {  // it give by default only key
		fmt.Println(v)
	}


	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte, 2 byte ... due to unicode not ascii value
	for i, c := range "Anil" {
		fmt.Println(i, c)
	}

	for _, c := range "Anil" {
		fmt.Println(c)
	}

	for i:= range "Anil" {
		fmt.Println(i)
	}
}
