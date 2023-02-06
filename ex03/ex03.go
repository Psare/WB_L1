package main

import (
	"fmt"
	"sync"
)

var b int // создаю глобалку в который будет хранится сумма

func square(mas []int, i int) int {
	b = mas[i] * mas[i] + b
	return b
}

func main() {
	var wg sync.WaitGroup
	var mas = []int{2, 4, 6, 8, 10}
	// Пока не прочитает всё из канала
	for i := range mas {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			square(mas, i)
		}(i)
	}
	wg.Wait()
	fmt.Println(b)
}