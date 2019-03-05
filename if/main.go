// Package main calculate square root of given integer
package main

import (
	"fmt"
	"go-tour/if/lib"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main() {
	fmt.Println(sqrt(4), sqrt(-4))  // 2 2i
	fmt.Println(lib.Pow(2, 5, 100)) // 32
	fmt.Println(lib.Pow(2, 5, 10))  // 10

	fmt.Println(
		lib.Pow2(3, 2, 10),
		lib.Pow2(3, 3, 20),
	) // 9 20

}
