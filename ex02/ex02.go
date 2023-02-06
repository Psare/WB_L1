package main

import (
	"fmt"
	"sync"
	"time"
)

// Принимает число "mas" и waitgroup wg, считает квадрат числа и выводит.
func print(mas int, wg *sync.WaitGroup) {
	fmt.Println(mas * mas)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mas = []int{2, 4, 6, 8, 10}
	// Пока не прочитает всё из канала
	for i := range mas {
		wg.Add(1)
		// вызов горутины
		go func(i int) {
			print(mas[i], &wg)
		}(i)
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
}