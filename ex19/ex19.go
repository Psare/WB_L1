package main

import (
	"fmt"
)
//идем по строке с начала и конца к середине, меняя местами символы
func reverse(s string) string {
	runes := []rune(s) //создаем руну
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	fmt.Println("Original string:", str)
	fmt.Println("Reversed string:", reverse(str))
}