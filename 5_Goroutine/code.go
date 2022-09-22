package main

import (
	"fmt"
	"time"
)

// Goroutines are used to implement concurrency in go
func main() {
	go greeter("hello")
	greeter("world")
}
func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}
