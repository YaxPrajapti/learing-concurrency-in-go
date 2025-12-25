package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(num int, msgChan chan string) {
	fmt.Printf("Sending some message %d\n", num)

	time.Sleep(time.Second * time.Duration(num))

	msg := fmt.Sprintf("Message %d sent", num)

	msgChan <- msg // sending message to the channel
}

func recieveMessage(msgs <-chan string) {
	fmt.Println("Waiting for messages")
	// this is for loop over channel which means, keep recieving until channel is closed
	for msg := range msgs {
		fmt.Println("Recieved: ", msg)
	}
}

// channels are used to communicate between routines
func main() {
	msgChan := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		sendMessage(2, msgChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sendMessage(3, msgChan)
	}()

	/*
		fmt.Println(<-msgChan) // reading from channel is blocking operation
		fmt.Println(<-msgChan)
	*/

	go func() {
		wg.Wait()
		close(msgChan) // this will close the channel and make the resources free
	}()
	recieveMessage(msgChan)
	fmt.Println("Work done")
}
