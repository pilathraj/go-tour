// Package lib - calculate pow
package lib

import (
	"fmt"
	"math"
)

// Pow2 calculate power
func Pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
		// same line.
	} else {
		fmt.Printf("%g >= %g\n-------\n", v, lim)
	}
	// can't use v here, though
	return lim
}
