package main

import "fmt"

var add = func(a int, b int) int {
	return a + b
}

func main() {
	var sub = func(a int, b int) int {
		return a - b
	}
	fmt.Println("5 + 3 =", add(5, 3))
	fmt.Println("5 - 3 =", sub(5, 3))

	immediate()
}
