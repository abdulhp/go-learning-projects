package main

import (
	"fmt"
)

func main() {

    messages := make(chan string)

    go func() { messages <- "ping" }()
	go func() { messages <- "ping2" }()

    msg := <-messages
    fmt.Println(msg)
	msg = <-messages
    fmt.Println(msg)
}