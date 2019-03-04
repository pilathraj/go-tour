// Package main for calculate sum of 0 to 9
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		//fmt.Println(i)
	}
	fmt.Println("Sum: ", sum) //Sum:  45
}
