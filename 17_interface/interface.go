package main

import "fmt"

type Paymentor interface{
	Pay(amount float32)
}


type PaymentGateway struct{
//    gateway Razorpay  // hardcoded
//    gateway2 Stripe   // hardcoded, to solve this problem, we need interface

	gateway3 Paymentor
}

type Razorpay struct{}

func (R Razorpay) Pay(amount float32){
	fmt.Println("making payment with razorpay", amount)
}

type Stripe struct{}

func (S Stripe) Pay(amount float32){
	fmt.Println("making payment with stripe", amount)
}


func (P PaymentGateway) MakePayment(amount float32){
	// P.gateway.pay(amount)
	// P.gateway2.pay(amount)

	P.gateway3.Pay(amount)
}

// in futute, if we ask to add another payment gateway. for ex..
type Paypal struct{}

func (P Paypal) Pay(amount float32){
	fmt.Println("making payment with PayPal ", amount)
}

func main() {
	// // rz := Razorpay{}
	// pg := PaymentGateway{
	// 	// gateway: rz,
	// 	gateway: Razorpay{},
	// 	gateway2: Stripe{},

	// }
    // pg.gateway.pay(20.0)
	// pg.gateway2.pay(39.3)



	// Scenario 1: Using PayPal
	myPaypal := Paypal{}
	
	pg := PaymentGateway{
		gateway3: myPaypal, // "Plugging in" Paypal
	}

	pg.MakePayment(40)

	// Scenario 2: Switching to Stripe is easy now!
	pg2 := PaymentGateway{
		gateway3: Stripe{}, // "Plugging in" Stripe
	}
	pg2.MakePayment(100)
	
}