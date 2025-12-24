package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i :=5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Others")
	}


	// multiple condition switch

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday : 
		fmt.Println("It's party time")
	default:
		fmt.Println("Its a working day")
	}

	// type switch
	whoAmI := func(i interface{})  {
		switch i.(type){
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("other type")
		}
	}

	whoAmI("goolang")
	whoAmI(49)

}
