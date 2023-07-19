package main

import "fmt"

func main() {
	// how to create an array
	a := [3]string{"A", "B", "C"}

	fmt.Println(a[0]) //A
	fmt.Println(a[1]) //B
	fmt.Println(a[2]) //C

	a[1] = "b" // overwrite element

	fmt.Println(a[1]) // b

	fmt.Println("----------------")

	// how to create an array without give the number of elements
	b := [...]string{"X", "Y", "Z"}
	fmt.Println(b[1])
}
