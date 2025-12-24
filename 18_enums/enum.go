package main

import "fmt"

type OrderStatus int

// 1. Use iota for automatic numbering (0, 1, 2, 3...)
const (
	Confirmed OrderStatus = iota // 0
	Received                     // 1
	Prepared                     // 2
	Delivered                    // 3
)

// 2. Add a String() method to your type.
// fmt.Println automatically looks for this method to know how to print the value!
func (o OrderStatus) String() string {
	// Simple slice/array lookup
	statuses := []string{"Confirmed", "Received", "Prepared", "Delivered"}
	
	// Safety check: ensure index is valid
	if o < 0 || int(o) >= len(statuses) {
		return "Unknown"
	}
	return statuses[o]
}

func ChangeOrderStatus(currentStatus, newStatus OrderStatus) {
    // 3. Using Printf (F-string alternative)
	fmt.Printf("Changing order status from %s to %s\n", currentStatus, newStatus)
}

func main() {
	// Internally these are numbers (0, 1), but they print as strings!
	ChangeOrderStatus(Confirmed, Received)
	ChangeOrderStatus(Prepared, Delivered)
}