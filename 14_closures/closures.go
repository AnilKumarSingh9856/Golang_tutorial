package main

import "fmt"

func counter() func() int {
	count :=0

	return func() int {
		count +=1
		return  count
	}
}


func main(){
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	

}

/*
1. What is a Closure?
A closure is a function value that references variables from outside its own body. The function may access and assign to the referenced variables.

In our code:

The anonymous function func() int is the Closure.

It "closes over" (captures) the variable count.

Even though the outer function counter() finishes and returns, the variable count stays alive in memory because the inner function still needs it.
*/
