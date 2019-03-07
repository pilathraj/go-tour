// Package main pointer
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         //  point to i.
	fmt.Println(p)  // memory address of i. (eg. 0xc000064078)
	fmt.Println(*p) // read i through the pointer 42
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i is 21

	p = &j          // point to j
	fmt.Println(p)  // memory address of i. (eg. 0xc000064080)
	fmt.Println(*p) // read j through the pointer 2701
	*p = *p / 37    // divide j through the pointer
	fmt.Println(j)  // see the new value of j is 73
}
