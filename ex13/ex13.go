package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println("Изначальные данные : a =", a,"b =", b)
	a = a ^ b //11110
	b = a ^ b //01010
	a = a ^ b //10100
	fmt.Println("a =", a,"b =", b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a =", a, "b =", b)
	a, b = b, a
	fmt.Println("a =", a, "b =", b)
}