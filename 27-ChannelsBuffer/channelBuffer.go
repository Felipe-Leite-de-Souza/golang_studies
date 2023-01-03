package main

import (
	"fmt"
)

func main() {
	channel := make(chan string, 2)
	channel <- "Hello World!"
	channel <- "Hello World1!"

	message := <-channel
	message2 := <-channel
	fmt.Println(message)
	fmt.Println(message2)
}
