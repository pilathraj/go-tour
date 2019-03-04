// Package main for calculate sum of 0 to 9
package main

import (
	"fmt"
	"go-tour/for-loop/lib"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		//cd .fmt.Println(i)
	}
	fmt.Println("Sum: ", sum) //Sum:  45
	sum = lib.Add()
	fmt.Println("Sum: ", sum) //Sum:  1024
	sum = lib.GosWhile()
	fmt.Println("Sum: ", sum) //Sum:  128

}
