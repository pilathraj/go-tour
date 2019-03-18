package main

import "fmt"

func makingSlice() {
	a := make([]int, 5)
	printSlice2("a:", a) //a: len=5, cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5)
	printSlice2("b:", b) //b: len=0, cap=5 []

	c := b[:2]
	printSlice2("c:", c) //c: len=0, cap=5 [0,0]

	d := c[1:5]
	printSlice2("d:", d) //d: len=4, cap=4 [0,0,0,0]

}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
