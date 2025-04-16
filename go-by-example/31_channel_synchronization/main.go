// Program for channel synchronization
package main

import (
	"fmt"
	"time"
)

func messageChannel(controlChannel chan bool) {
	messages := make(chan string)
	go func() { time.Sleep(5 * time.Second); messages <- "ping" }()
	go func() { messages <- "ping2" }()

	// Block until we receive a message from the channel
	msg := <-messages
	fmt.Println(msg)

	// Block until we receive another message from the channel
	msg = <-messages
	fmt.Println(msg)

	controlChannel <- true
}

func main() {
	controlChannel := make(chan bool, 1)

	go messageChannel(controlChannel)

	if <-controlChannel == false {
		fmt.Println("Channel closed unexpectedly")
	}
	fmt.Println("All messages received")
}
