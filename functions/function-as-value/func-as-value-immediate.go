package main

import "fmt"

var add1 = func(a int, b int) int {
	return a + b
}(5, 3)

func immediate() {
	var sub = func(a int, b int) int {
		return a - b
	}(5, 3)
	fmt.Println("5 + 3 =", add1)
	fmt.Println("5 - 3 =", sub)
}
