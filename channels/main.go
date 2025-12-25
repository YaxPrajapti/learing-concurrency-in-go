package main

import (
	"fmt"
	"time"
)

func sendMessage(num int, msgChan chan string) {
	fmt.Printf("Sending some message %d\n", num)

	time.Sleep(time.Second * time.Duration(num))

	msg := fmt.Sprintf("Message %d sent", num)

	msgChan <- msg // sending message to the channel
}

// channels are used to communicate between routines
func main() {
	msgChan := make(chan string)
	go sendMessage(2, msgChan)
	go sendMessage(3, msgChan)

	fmt.Println(<-msgChan) // reading from channel is blocking operation
	fmt.Println(<-msgChan)

	close(msgChan) // this will close the channel and make the resources free
}
