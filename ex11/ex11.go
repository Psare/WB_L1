package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 7, 4, 9}
	set2 := []int{4, 5, 6, 7, 1, 9, 3}

	// Создаем карту для хранения элементов множества set1
	setMap := make(map[int]bool)
	for _, value := range set1 {
		setMap[value] = true
	}

	intersection := make([]int, 0)

	// Проверяем, присутствует ли каждый элемент set2 в setMap
	for _, value := range set2 {
		if setMap[value] {
			intersection = append(intersection, value)
		}
	}

	fmt.Println("Intersection:", intersection)
}