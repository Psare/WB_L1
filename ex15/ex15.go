package main

import "fmt"
//Символ могут занимать разное кол-во байт, поэтому нужно использовать руну
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
  fmt.Println("orig", justString)
  r := []rune(v)
  justStringRune := string(r[:100])
  fmt.Println("rune", justStringRune)
}

func createHugeString(size int64) string{
	var s string
	var i int64
	for i = 0; i < size; i++ {
		s += "0"
	}
	return s
}

func main() {
  someFunc()
}
