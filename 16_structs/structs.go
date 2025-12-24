package main

import (
	"fmt"
	"time"
)


type Customer struct{
	name, phone string
}

// Standard: Capitalize fields to make them "Exported" (public)
type Order struct {
	ID         string
	Amount     float32
	Status     string
	ReceivedAt time.Time 
	Customer    // struct embedding
}

// Convention: Rename 'new' to 'NewOrder' to avoid shadowing built-in 'new'
func NewOrder(id string, amount float32, status string, receivedAt time.Time) *Order {
	// You can return the address directly
	return &Order{
		ID:         id,
		Amount:     amount,
		Status:     status,
		ReceivedAt: receivedAt,
	}
}

// Receiver name 'o' is good. Use CamelCase for arguments.
func (o *Order) ChangeAmount(newAmount float32) {
	o.Amount = newAmount
}

func (o Order) GetAmount() float32 {
	return o.Amount
}

func main() {
	// 1. Using Struct Literal
	myOrder := Order{
		ID:     "1",
		Amount: 50.00,
		Status: "received",
	}
	myOrder.ReceivedAt = time.Now()

	fmt.Printf("Order 1: %+v\n", myOrder) // %+v prints field names too!

	// 2. Using the Constructor (You defined it, so let's use it!)
	myOrder2 := NewOrder("2", 299, "paid", time.Now())
	
	myOrder2.ChangeAmount(200)
	
	// Since myOrder2 is a pointer (from NewOrder), we dereference it with * to print value
	fmt.Printf("Order 2: %+v\n", *myOrder2)

	// Anonymous struct
	lang := struct {
		Name   string
		IsGood bool
	}{"GoLang", true}

	fmt.Println(lang)
	
	// custs := Customer{
	// 	name: "Anil",
	// 	phone: "0123456789",
	// }

	myOrder4 := Order{
		ID: "6",
		Amount: 200,
		Status: "got it",
		// Customer: custs,
		Customer: Customer{
			name: "Dev",
			phone: "9876543210",
		},
	}
    fmt.Println(myOrder4.Customer)
    fmt.Println(myOrder4)
}