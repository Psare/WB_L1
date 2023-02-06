package main

import (
	"fmt"
	"strings"
)

// Принимает на вход строку str, возвращает true, если уникальные
func check(str string) bool {
	// Сначала переводим все в нижний регистр
	s := strings.ToLower(str)
	m := make(map[string]int)
	// Если находит одинаковые то false
	for _, value := range s {
		if _, ok := m[string(value)]; ok {
			return false
		}
		m[string(value)]++
	}
	return true
}

func main() {
	str1 := "abcde"
	str2 := "abCdefA"
	str3 := "jcvnkevwjvn"
	fmt.Println(str1,":", check(str1))
	fmt.Println(str2,":", check(str2))
	fmt.Println(str3,":", check(str3))
}