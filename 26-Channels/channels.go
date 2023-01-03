package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go escrever("Hello World!", channel)

	fmt.Println("Depois da função....")

	//for {
	//	message, channelOpen := <-channel
	//	if !channelOpen {
	//		break
	//	}
	//	fmt.Println(message)
	//}
	for message := range channel {
		fmt.Println(message)
	}
	fmt.Println("Fim do programa!")
}

func escrever(texto string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- texto
		time.Sleep(time.Second)
	}
	close(channel)
}
