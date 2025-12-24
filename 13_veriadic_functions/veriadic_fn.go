package main

import "fmt"

func sum(nums ...int) int {
	sum :=0
	for num := range nums{
		sum +=num
	}
	return  sum
}

func printAnything(value ...interface{}){
	for _, val := range value{
		fmt.Print(val, " ")
	}
}

func main() {
	
	fmt.Println(1,2,3,4,"Hello", true) // veriadic fn

	result := sum(1,2,3,4,5)
	fmt.Println(result)
	slices := []int{1,2,3,4,5,6,7}
	fmt.Println(sum(slices...))

	printAnything("Anil", 2, 5, true, "Kumar")
}
