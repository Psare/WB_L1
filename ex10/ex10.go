package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5}
	groups := map[int][]float64{}
	//Заолняем значениями
	for _, temp := range temperatures {
		group := int(math.Floor(temp / 10)) * 10
		groups[group] = append(groups[group], temp)
	}
	//разбиваем мапу на группы
	for group, temps := range groups {
		merged := fmt.Sprintf("%d: ", group)//после группы идет двоеточие
		//и температуры
		for _, temp := range temps {
			merged += fmt.Sprintf("%.1f, ", temp)//после температуры - запятая пробел
		}
		fmt.Println(merged)//выводим
	}
}