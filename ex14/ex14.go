package main

import (
    "fmt"
)

func opredelit(data interface{}){
	switch data.(type) {
    case int:
        fmt.Println("Integer")
    case string:
        fmt.Println("String")
    case bool:
        fmt.Println("Boolean")
    case chan int:
        fmt.Println("Channel")
    default:
        fmt.Println("Unknown Type")
    }
}

func main() {
    var data interface{}
    data = 10
    opredelit(data)
	data = 10.0
    opredelit(data)
	data = "hello!"
    opredelit(data)
	data = true
    opredelit(data)
	ch := make(chan int)
	data = ch
    opredelit(data)
}