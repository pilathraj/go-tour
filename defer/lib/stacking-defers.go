// Package lib stack defer
package lib

import "fmt"

func StackDefer() {
	fmt.Print("Start ")
	i := 0

	for i < 10 {
		i = i + 1
		defer fmt.Print(i, " ")
	}
	fmt.Print("done! ")
}
