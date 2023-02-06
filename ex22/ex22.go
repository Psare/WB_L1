package main

import (
	"fmt"
	"math/big"
)

func main() {
  a := big.NewInt(1 << 30) // можно использовать int64
  b := big.NewInt(2 << 30)
  c := big.NewInt(0)

  fmt.Println("a:", a)
  fmt.Println("b:", b)
  fmt.Println("a + b:", c.Add(a, b))
  fmt.Println("a - b:", c.Sub(a, b))
  fmt.Println("a * b:", c.Mul(a, b))
  fmt.Println("a / b:", c.Div(a, b))
}