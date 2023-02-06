package main

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем карту для хранения уникальных элементов 
	set := make(map[string]bool)
	for _, s := range sequence {
		set[s] = true
	}

	// Извлекаем уникальные элементы из карты и сохраняем их в срезе
	uniqueElements := make([]string, 0, len(set))
	for k := range set {
		uniqueElements = append(uniqueElements, k)
	}

	fmt.Println("Unique elements of the sequence:", uniqueElements)
}