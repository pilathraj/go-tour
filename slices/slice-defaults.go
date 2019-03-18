package main

import "fmt"

// SliceDefaults - example for Slice Defaults
func SliceDefaults() {
	s := []int{1, 2, 3, 4, 5, 1}

	s = s[1:4]     //[2,3,4]
	fmt.Println(s) //[2,3,4]

	s = s[:2] //[2, 3]
	fmt.Println(s)

	s = s[1:] //[3]
	fmt.Println(s)

}
