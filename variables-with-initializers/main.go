// Package main short variable declarations
package main

import "fmt"

var i, j int = 1, 2 // variable declaration at package level

func main() {
	var c, python, java = true, false, "Yes" //  variable will take the type of the initializer
	fmt.Println(i, c, python, java)          // output:  1 true false Yes
}
