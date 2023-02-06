package main

import "fmt"

func removeElement(s []int, i int) []int {
  s = append(s[:i], s[i+1:]...)
  return s
}

func main() {
  slice := []int{1, 2, 3, 4, 5}
  fmt.Println("Original slice:", slice)

  slice = removeElement(slice, 3)
  fmt.Println("Slice after removing element:", slice)
}