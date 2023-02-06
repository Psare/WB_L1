package main

import (
	"fmt"
	"time"
)

func sendValues(c chan int) {
	for i := 0; i < 100000; i++ {
		c <- i
		time.Sleep(200 * time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan int)
	go sendValues(c)//отправляем данные в канал

	timeout := time.After(3 * time.Second)
	for {
		select {
		case value, ok := <-c:
			if !ok {
				fmt.Println("Channel closed.") // если значения закнчились
				return
			}
			fmt.Println("Received value:", value) // вывод
		case <-timeout: //если таймаут
			fmt.Println("Timeout.")
			return
		}
	}
}