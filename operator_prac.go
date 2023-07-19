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

	// same or not
	fmt.Println(x == y) // false
	fmt.Println(x != y) // true

	// and, or
	fmt.Println(x >= 6 && y <= 4) // true
	fmt.Println(x <= 6 || y <= 4) // true

	// Compound assignment operators
	x += y
	println(x) // 13

}
