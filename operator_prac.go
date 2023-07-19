package main

import "fmt"

func main() {
	// +,-,*,/,%
	x := 10
	y := 3

	fmt.Println(x + y) // 13
	fmt.Println(x - y) // 7
	fmt.Println(x * y) // 30
	fmt.Println(x / y) // 3
	fmt.Println(x % y) // 1

	// relationship
	fmt.Println(x > y) // true
	fmt.Println(x < y) // false

}
