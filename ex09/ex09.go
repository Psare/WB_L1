package main

import "fmt"

// Принимает массив и канал, в который будет записывать числа из массива.
func chanel1(arr []int, ch1 chan int) {
	defer close(ch1)
	for _, ch := range arr {
		ch1 <- ch
	}
}

// Принимает данные из канала ch1 и записывает в канал ch2.
func chanel2(ch2 chan int, ch1 chan int) {
	defer close(ch2)
	for ch := range ch1 {
		ch2 <- ch * 2
	}
}

func main() {
	// Создаем каналы ch1 и ch2 и массив, из которого будем брать числа.
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	// Запускаем горутины и печатаем содержимое второго канала.
	go chanel1(arr, ch1)
	go chanel2(ch2, ch1)

	for ch := range ch2 {
		fmt.Println(ch)
	}
}