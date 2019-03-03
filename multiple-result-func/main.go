// Package main swap function used to swap the strings
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	fmt.Println(swap("Hello", "Pilathraj"))
}
