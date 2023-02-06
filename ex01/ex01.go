package main

import "fmt"

//создаю структуру родителя
type Human struct {
	Name string
}

//создаю структуру ребенка, в котором встроен родитель
type Action struct {
	Human
	Age int
}

func main() {
	var John Action

	John.Name = "John" //имя из хумана
	John.Age = 29 // возраст из ребенка
	fmt.Println(John)
	fmt.Printf("My name is %s and i am %d years old\n", John.Name, John.Age)
}