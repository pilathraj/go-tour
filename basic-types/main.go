// Package main type and value
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // Type: bool Value: true
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", z, z)           // Type: complex128 Value: (2+3i)
}
