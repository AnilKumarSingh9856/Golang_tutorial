package main

import (
	"fmt"
	"time"
)

// receiving the value
func receive(num chan int, num1, num2 int){
	sum :=  num1 + num2
	num <- sum
}

// sending the value
func send(nums chan int){
	for i := range nums{
		fmt.Println("Processing the value of ", i)
		time.Sleep(time.Microsecond)
	}
	
}

func task(done chan bool){
	defer func ()  {done <- true}()

	fmt.Println("processing...")
}

func processIt(value chan string){
	for i:=0; i<10; i++{
		value <- fmt.Sprintf("%d@gmail.com",i)
		
	}
	close(value) // sending signal, i am done sending
}

func main(){
    // res := make(chan string, 10)

	// go processIt(res)

	// for email := range res {
	// 	fmt.Println("Received email:", email)
	// }

	// res <- "1@gmail.com"
	// res <- "2@gmail.com"

	// fmt.Println(<- res)
	// fmt.Println(<- res)




	// unbuffer channels

	// istrue := make(chan bool)
	
	// go task(istrue)

	// <- istrue  // block

	// rv := make(chan int)
    // go send(rv)

	// for {
	// 	rv <- rand.Intn(10)
	// }


	// intChan := make(chan int)

	// go receive(intChan, 4, 5)

	// fmt.Println(<- intChan)


	// strChan := make(chan string)   // No capacity defined

	// // Deadlock!
	// // The main thread waits for someone to take the message
	// // But no one is there. The program crashes.

	// strChan <- "Anil"

	// res := <- strChan

	// fmt.Println(res) // blocking bec it's size is empty

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func ()  {
		chan1 <- 5
	}()

	go func ()  {
		chan2 <- "Surjit"
	}()

	for i:=0; i<2; i++{
		select{
		case chan1Value := <- chan1 :
			fmt.Println("received data from chan1", chan1Value)
		case chan2Value := <- chan2:
			fmt.Println("received data from chan2", chan2Value)
		}
	}

		
}
