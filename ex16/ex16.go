package main

import (
	"fmt"
	"sort"
)
//partion - функция которая принимает фрагмент целых чисел,
// младший индекс и старший индекс и разбивает срез так, что
// элементы, меньшие или равные стержню, находятся слева от стержня
// а элементы, превышающие опорную точку, находятся справа от нее
func partition(numbers []int, low int, high int) int {
	// Выбираем последний элемент в срезе
	pivot := numbers[high]
	i := low - 1
	// срез от low до high-1
	for j := low; j < high; j++ {
		// Если текущий элемент меньше или равен pivot
		// увеличиваем i и меняем местами текущий элемент с элементом на i
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	// Поменяем pivot с элементом по адресу i+1
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	// Возвращаем индекс
	return i + 1
}

func quicksort(numbers []int) {
	quicksortHelper(numbers, 0, len(numbers)-1)
}

func quicksortHelper(numbers []int, low int, high int) {
	if low < high {
		p := partition(numbers, low, high)
		// Рекурсивно отсортировать подмассив от low до p-1
		quicksortHelper(numbers, low, p-1)
		// Рекурсивная сортировка подмассива от p+1 до high
		quicksortHelper(numbers, p+1, high)
	}
}

func main() {
	numbers := []int{5, 2, 9, 1, 5, 6}
	numbers1 := []int{5, 2, 9, 1, 5, 6}
	sort.Ints(numbers)
	fmt.Println(numbers)
	quicksort(numbers1)
	fmt.Println(numbers1)
}