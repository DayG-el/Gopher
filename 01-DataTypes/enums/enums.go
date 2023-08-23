package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct for Order
type Order struct {
	Id     int
	Price  float64
	Status OrderStatus
}

// Define a custom type for order status
type OrderStatus int

// Define constants for the order status
const (
	Created        = iota //0
	Processing     = iota //1
	WaitForPayment = iota //2
	PaymentDone    = iota //3
	Issue          = iota //4
	Refund         = iota //5
)

func main() {
	order1 := Order{Id: 1, Price: 100.0, Status: Created}
	order2 := Order{Id: 2, Price: 200.0, Status: PaymentDone}
	order3 := Order{Id: 3, Price: 300.0, Status: Refund}
	order1Json, _ := json.Marshal(order1)
	order2Json, _ := json.Marshal(order2)
	order3Json, _ := json.Marshal(order3)
	println(string(order1Json))
	println(string(order2Json))
	println(string(order3Json))
	fmt.Printf("%+v", order3)
}
