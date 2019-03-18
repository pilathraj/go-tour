package main

import "fmt"

func nilSlices() {
	var s []int
	printSlice(s)
	if s == nil {
		fmt.Println("nil!")
	}
}
