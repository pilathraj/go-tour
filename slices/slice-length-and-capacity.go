package main

import "fmt"

//SliceLenCaps - example for Slice length and capacity
func SliceLenCaps() {
	s := []int{1, 2, 3, 4, 5, 7, 9, 0}
	printSlice(s) //len=8, cap=8 [1 2 3 4 5 7 9 0]

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) //len=0, cap=8 []

	// Extend its length.
	s = s[:4]
	printSlice(s) //len=4, cap=8 [1 2 3 4]
	// Drop its first two values.
	s = s[2:]     // look capacity reduced.
	printSlice(s) //len=2, cap=6 [3 4]
	//s = s[4:]     // slice bounds out of range error
	//printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}
