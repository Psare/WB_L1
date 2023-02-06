package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
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

	// Запрашиваем у пользователя количество рабочих
	fmt.Print("Enter the number of workers: ")
	input := bufio.NewReader(os.Stdin)
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	n, _ := strconv.Atoi(text)
	if n == 0 {
		os.Exit(1)
	}
	// Создаем wait group и ждем финиша всех workers
	var wg sync.WaitGroup
	wg.Add(n)

	// начало n workers
	for i := 1; i <= n; i++ {
		go worker(i, &wg)
	}

	// Запись данных в канал
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