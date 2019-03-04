// Package main variable declaration with initializer
package main

import "fmt"

func main() {
	var i, j int = 1, 2                   //  type with var declaration
	k := 3                                // Short variable declarations with implicit type
	c, python, java := true, false, "no!" //   := construct is available only inside the func
	fmt.Println(i, j, k, c, python, java) // output:  1 2 3 true false no!
}
