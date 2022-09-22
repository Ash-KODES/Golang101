package main

import "fmt"

func main() {
	//"_"-->Blank identifier
	//sum := 0
	var a []int
	for i := 0; i < 5; i++ {
		var value int
		fmt.Scan(value)
		a = append(a, value)
	}
	fmt.Println(a)
}
