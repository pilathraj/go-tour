// package main - print current os name
package main

import (
	"fmt"
	"go-tour/defer/lib"
)

func main() {
	defer fmt.Println("Pilathraj!")
	fmt.Println("Hello ")
	lib.StackDefer()
}

// op1: Hello
// Pilathraj!

//0p2:
//Hello
//Start done! 10 9 8 7 6 5 4 3 2 1 Pilathraj!
