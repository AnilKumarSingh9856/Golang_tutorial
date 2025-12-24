package main

import "fmt"

func main() {
	age :=20

	if age < 18 {   // other way of writing ->  if age :=20; age >= 18  
		fmt.Println("Teenager")
	} else if age < 50 {
		fmt.Println("You are adult person")
	}else {
		fmt.Println("You are a senior citizen")
	}


	// go does not have tenary, you will have to use normal if_else
}
