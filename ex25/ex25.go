package main

import (
	"fmt"
	"time"
)

func Sleep(duration int) {
	<-time.After(time.Duration(duration) * time.Second)
}

func main() {
	fmt.Println("Started sleeping")
	Sleep(2)
	fmt.Println("Finished sleeping")
}