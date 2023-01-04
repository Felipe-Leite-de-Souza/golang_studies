package main

import (
	"fmt"
	"time"
)

func main() {
	channel := multiplexar(escrever("Hello World!"), escrever("Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplexar(inputChannel1, inputChannel2 <-chan string) <-chan string {
	outputChannel := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChannel1:
				outputChannel <- message
			case message := <-inputChannel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func escrever(texto string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * time.Duration(500))
		}
	}()

	return channel

}
