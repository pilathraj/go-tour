// Package main - arrays

package main

import "fmt"

func main() {
	var a [10]int
	a[0] = 1
	a[2] = 2
	a[3] = 3
	//a[11] = 11 // array index 11 out of bouunds[0:6]
	fmt.Println(a)
	fmt.Println(a[3]) // 3

	prime := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(prime)

}
