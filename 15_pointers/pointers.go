package main
import "fmt"

func fn1(num int) int {
    num = 3         // This modifies the COPY
    return num      // Returns 3
}

func fn2(num *int) int{
   *num = 4         // This modifies the ORIGINAL
   return  *num
}

func main() {
    num := 5
    
    // TEST 1: Pass by Value
    fmt.Println("Function Returns:", fn1(num)) // Prints 3
    fmt.Println("Value of num:", num)          // Prints 5 (STILL 5!)
    
    fmt.Println("----------------")

    // TEST 2: Pass by Reference
    fmt.Println("Function Returns:", fn2(&num)) // Prints 4
    fmt.Println("Value of num:", num)           // Prints 4 (CHANGED!)
}