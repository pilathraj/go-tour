// Package main - slice example
package main

import "fmt"

func main() {
	//slice
	prime := [7]int{1, 2, 3, 4, 5, 6, 7}
	s1 := prime[1:4]
	fmt.Println(s1) // [2,3,4]

	names := [6]string{
		"Raja",
		"Raj",
		"Ragu",
		"Ravi",
		"John",
		"Paul",
	}
	fmt.Println(names) //[Raja,Raj,Ragu,Ravi, John,Paul]
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) //[Raja,Raj][Raj,Ragu]

	b[0] = "XXX"       // it replace Raj
	fmt.Println(a, b)  //[Raja,XXX][XXX,Ragu]
	fmt.Println(names) //[Raja,XXX,Ragu,Ravi, John,Paul]
	SliceLiterals()    // same package so no need to dot no name here :)
	SliceDefaults()
	SliceLenCaps()
	nilSlices() // Note: samoe package No need to UpperCase func name :)
	makingSlice()
	slicesOfSlices()
	appendInfo()
	rangeInfo()
}
