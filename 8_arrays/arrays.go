package main

import "fmt"

// numbered sequence of specific length

func main() {
	var arr [4]int   // by default its value is [0, 0,...., 0]
	arr[1] = 4
	var arr2 [5]string // by default its value is [ ] (empty)

	var arr3 [4]bool // by default its value is [false,...]
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)

	// 2d arrays

	nums := [2][2]int{{}, {}}
	fmt.Println(nums)



	// fixed size, that is predictable
	// memory optimization
	// constant time access
}
