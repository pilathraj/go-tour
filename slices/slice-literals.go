package main

import "fmt"

// SliceLiterals - slice-literals example
func SliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13} // no length  --- so it a slice
	fmt.Println(q)                 //[2, 3, 5, 7, 11, 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) //[true, false, true, true, false, true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) //[{2, true},{3, false}, {5,true}, {7,true},{11, false}, {13,true}]

}
