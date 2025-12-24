package main

import (
	"fmt"
)

// most used construct in go

// In Go, there are several ways to create a slice:

// Using the []datatype{values} format -> slice_name := []datatype{values}

// Create a slice from an array  -> var myarray = [length]datatype{values} // An array
// myslice := myarray[start:end] // A slice made from the array

// Using the make() function  -> slice_name := make([]type, length, capacity)

func main() {
	myslice := []int{}
	fmt.Println(myslice==nil)
	fmt.Println(myslice)

   // initialize the slice during declaration
	myslice1 := []int{1,2,3}
	myslice1[2] = 5
	// adding new element
	myslice1 = append(myslice1, 7, 8,9)
	fmt.Println(myslice1)
	

	arr1 := [6]int{10, 11, 12, 13, 14,15}
  	myslice2 := arr1[2:4]
	fmt.Println(myslice2)  // a slice made from the array

	// appending one slice to another
	new_slices := append(myslice1, myslice2...)
	fmt.Println(new_slices)

    // Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	// Note: If the capacity parameter is not defined, it will be equal to length.

	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))
}


// In Go, there are two functions that can be used to return the length and capacity of a slice:

// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
