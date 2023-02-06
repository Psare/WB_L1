package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("goroutine stopped")
				return
			default:
				fmt.Println("goroutine running")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	close(stop)
	time.Sleep(2 * time.Second)
}