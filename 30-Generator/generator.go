package main

import (
	"fmt"
	"time"
)

func main() {
	channel := escrever("Hello World!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func escrever(texto string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel

}
