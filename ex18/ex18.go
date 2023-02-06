package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}
    fmt.Println("Preprogram Counter Valut", counter.Value())
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter Value:", counter.Value())
}