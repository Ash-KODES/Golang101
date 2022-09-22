package main

import "fmt"

// Defer-->end mai execute karta hai
// if more than one defer then LIFO order is followed-->last deferred first executed

func main() {

	defer fmt.Println("world")

	defer fmt.Println("one")

	defer fmt.Println("two")

	fmt.Println("hello")

	//Prints number from 1 to k
	Num_print(5)
}
func Num_print(k int) {

	for i := 1; i <= k; i++ {

		defer fmt.Println(i)
	}
}
