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
		// The `go` keyword starts a new goroutine for each function call.
		// A goroutine is a lightweight, concurrent unit of execution managed by the Go runtime.
		// Goroutines are not OS threads; many goroutines are multiplexed onto a smaller number of OS threads.
		// The Go scheduler decides when and on which OS thread a goroutine runs,
		// and goroutines can run in parallel across multiple CPU cores.
		go processOrder(order)
		// Starting a goroutine does not block the main function.
		// If main() returns, the program exits immediately,
		// and any running goroutines are terminated.
		// Therefore, the main goroutine must explicitly wait
		// for other goroutines to finish (e.g., using sync.WaitGroup).
	}

}
