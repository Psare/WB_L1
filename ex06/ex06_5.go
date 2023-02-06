package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

var data = make(chan string)

// worker считывает данные из канала и записывает их в stdout
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Printf("Worker %d received: %s\n", id, d)
	}
}

func main() {
	// Создаем канал для приема сигналов прерывания (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	n :=3
	// Создаем wait group и ждем финиша всех workers
	var wg sync.WaitGroup
	wg.Add(n)

	// начало n workers
	for i := 1; i <= n; i++ {
		go worker(i, &wg)
	}
	go func() {
		i := 1
		for {
			data <- strconv.Itoa(i)
			i++
			time.Sleep(150 * time.Millisecond)
		}
	}()

	// Ждем сигнала прерывания (Ctrl+C)
	<-c

	// Закрываем канал данных для остановки всех workers
	close(data)

	// Ждем финиша всех workers
	wg.Wait()

	fmt.Println("\nProgram completed.")
}