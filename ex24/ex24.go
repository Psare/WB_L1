package main

import (
	"fmt"
	"math"
)
// Структура точек
type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}
// Функция подсчета дистанции между точками
func (p Point) Distance(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(1, 3)
	p2 := NewPoint(4, 5)
	distance := p1.Distance(p2)
	fmt.Println("Distance between points:", distance)
}