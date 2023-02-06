package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "snow dog sun"
	words := strings.Split(input, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	output := strings.Join(words, " ")
	fmt.Println(output)
}