package main

import (
	"fmt"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(order Order) {
	fmt.Printf("Preparing order for table number: %d...\n", order.TableNumber)

	// add delay
	time.Sleep(order.PrepTime)

	fmt.Printf("Order ready for table number: %d...\n", order.TableNumber)
}

func main() {
	// this block of code will display basic synchronous calling of function.
	orders := []Order{
		{TableNumber: 1, PrepTime: 1 * time.Second},
		{TableNumber: 2, PrepTime: 2 * time.Second},
		{TableNumber: 3, PrepTime: 3 * time.Second},
		{TableNumber: 4, PrepTime: 4 * time.Second},
		{TableNumber: 5, PrepTime: 5 * time.Second},
	}

	for _, order := range orders {
		processOrder(order)
	}
}
