package main

import (
	"fmt"
	"sort"
)

func binarySearch(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1
	
	for left <= right {
		middle := (left + right) / 2
		if numbers[middle] > target {
			right = middle - 1
		} else if numbers[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	
	return -1
}

func main() {
	numbers := []int{5, 2, 9, 1, 5, 6, 7}
	sort.Ints(numbers) // перед отправкй массив нужно отсортировать
	target := 9
	index := binarySearch(numbers, target)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index of", target, "is", index)
	}
}