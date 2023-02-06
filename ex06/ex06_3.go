package main

import (
	"fmt"
	"time"
)

var stop bool

func main() {
	stop = false
	go func() {
		for {
			if stop {
				fmt.Println("goroutine stopped")
				return
			}
			fmt.Println("goroutine running")
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
	stop = true
	time.Sleep(2 * time.Second)
}