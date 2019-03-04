// Package main for constants
package main

import (
	"fmt"
)

const Pi = 3.14

func main() {
	const World = "Go"
	fmt.Println("Hello", World)      // Hello Go
	fmt.Println("Happy", Pi, "Days") // Happpy 3.14 Days
	const Truth = true
	fmt.Println("Go rules?", Truth) // Go rules true
}
