package main

import (
	"fmt"
	"sync"
)

// Принимает map m, данные i, WaitGroup wg, Mutex mu.
func write(m map[int]int, i int, wg *sync.WaitGroup, mu *sync.Mutex) {
	// Блокирование мютекса перед записью и уменьшаем счетчик wg перед завершением.
	mu.Lock()
	defer wg.Done()
	m[i] = i
	//Разблокировие мютекса.
	mu.Unlock()
}

func main() {
	// Создаем map m, waitgroup wg и mutex mu.
	m := make(map[int]int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 5; i++ { //Запускаем горутины.
		wg.Add(1)
		go write(m, i, &wg, &mu)
	}

	// Проверка завершения всех горутин.
	wg.Wait()
	fmt.Println(m)
}