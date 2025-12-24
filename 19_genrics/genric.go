package main

import (
	"fmt"
)

func printelement[T comparable](items []T){
	for _, val := range items{
		fmt.Println(val)
	}
}

// func printelement[T int | string](items []T){
// 	for _, val := range items{
// 		fmt.Println(val)
// 	}
// }



func printlIntSlices(arr []int){
	for _, val := range arr{
		fmt.Println(val)
	}
}
func printlStringSlices(arr []string){
	for _, val := range arr{
		fmt.Println(val)
	}
}


func main() {
	nums := []int{1,2,3,4}
	// printlIntSlices(nums)

	str := []string{"Anil", "Sunil", "Punil"}
	// printlStringSlices(str)
    printelement(nums)
    printelement(str)

}