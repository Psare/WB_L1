package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine stopped")
				return
			default:
				fmt.Println("goroutine running")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}