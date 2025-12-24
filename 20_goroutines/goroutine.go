package main

import (
	"fmt"
	"sync"
	"time"
)

// We create a "WaitGroup" to count how many tasks are running
var wg sync.WaitGroup

func makeCoffee() {
	// Schedule the "Done" signal for when this function finishes
	defer wg.Done() 
	
	fmt.Println("☕ Starting Coffee...")
	time.Sleep(2 * time.Second) // Simulate work taking time
	fmt.Println("☕ Coffee Ready!")
}

func fryEggs() {
	defer wg.Done()
	
	fmt.Println("🍳 Starting Eggs...")
	time.Sleep(2 * time.Second) // Simulate work taking time
	fmt.Println("🍳 Eggs Ready!")
}

func main() {
	fmt.Println("--- Breakfast Started ---")
	start := time.Now()

	// 1. Add 2 tasks to the counter
	wg.Add(2)

	// 2. Launch functions in the background using 'go'
	go makeCoffee()
	go fryEggs()

	// 3. Block here until the counter goes back to 0
	fmt.Println("... Waiting for breakfast ...")
	wg.Wait()

	fmt.Printf("--- Breakfast Finished in %v ---\n", time.Since(start))
}