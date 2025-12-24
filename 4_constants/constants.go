package main
import "fmt"

const learning = "Go Lang."
// const dt := 16   // we cannot define constant variable by using this := style.

func main()  {
	
	const lang  = "GoLang"

	// lang = "Go" // we will get error due to this line

	const yr = 2010
	fmt.Println(yr)

	const (
		port = 5000
		host = "localhost"
	)
	
}


/*
The Syntax Rules
For Variables: You can use var x = 1 (everywhere) or x := 1 (inside functions only).

For Constants: You must always use the = sign. The := operator is never allowed with const, regardless of whether you are inside or outside a function.

*/
