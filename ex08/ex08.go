package main

import "fmt"

func toOne(n *int64, i int) int64 {
	*n = *n | (1 << (i - 1))
	return *n
}

func toZero(n *int64, i int) int64 {
	*n = *n & ^(1 << (i - 1))
	return *n
}

func main() {
	var n int64
	n = 7
	bul := 1
	fmt.Printf("в двоичной:%b\n", n)
	if bul == 0 {
		fmt.Println(toZero(&n, 3))
	} else if bul == 1{
	fmt.Println(toOne(&n, 3))
	} else {
		fmt.Println("некоректный ввод")
	}
	fmt.Println(n)
}