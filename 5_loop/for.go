package main

import "fmt"

// for -> only construct in go for looping
func main() {
	
	// equilavent while loop
	i :=1
	for i<=4 {
		fmt.Println(i)
		i++   // without this line it is infinite while loop
	}

	// // infinite while loop

	// for {
	// 	println("1")
	// }


	// classic for loop

	for i:=1; i<=100; i++ {
		if i % 2 == 0 {
			continue
		}

		if i == 5 {
			break
		}

	}

	// same implemention with range keywords
	for i := range 11 {
		if i % 2 == 0 {
			continue
		}

		if i == 5 {
			break
		}
		fmt.Print("*", " ")
	}


	// for index, value := range array|slice|map {
	// // code to be executed for each iteration
	// }
}
