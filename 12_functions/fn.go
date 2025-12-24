package main

import "fmt"


func add(a int, b int) int {
	return a+b
}

func add1(a, b int) int {  // same as above parameter, bt shorthand
	return a+b
}

func return_multiple_value()(string, int, bool){ // returning multiple value
	return "Anil", 25, true
}

func processIt(fn func(a int) int){ // passing func as argument
	fn(1)
}

func processIt2() func(a int) int { // passing func as argument
	return func(a int) int  {
		return 4
	}
}



func main() {
	ans := add(3, 5)
	fmt.Println(ans)

	fmt.Println(return_multiple_value())
	val1, _, val3 := return_multiple_value()
	println(val1, val3)

	fn := func (a int) int {  // anonymous function
		return a
	}

	fmt.Println(fn(3))

	fn2 := processIt2()
	fmt.Println(fn2(3))

}
