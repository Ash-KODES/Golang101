package main

import "fmt"

func main() {

	var a []int
	for i := 0; i < 5; i++ {
		var value int
		fmt.Scan(&value)
		a = append(a, value)
	}
	fmt.Println(a)
}
