// Package main variable declaration
package main

import "fmt"

var c, python, java bool // variable declaration at package level

func main() {
	var i int                       // variable declaration at function level
	fmt.Println(i, c, python, java) // output:  0 false false false
}
