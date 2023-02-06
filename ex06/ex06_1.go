package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал done, чтобы сигнализировать о завершении функции
	done := make(chan bool)

	// Запускаем горутину, которая увеличивает счетчик каждую секунду
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine:", i)
			time.Sleep(time.Second)
		}

		// Сигнализируем о завершении выполнения программы, посылая значение в канал done
		done <- true
	}()

	// Ждем сигнала от выполнения программы, получая значение из канала done
	<-done
	fmt.Println("main function: done")
}