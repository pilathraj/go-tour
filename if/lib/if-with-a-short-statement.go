// Package main pow function return float64
package lib

import "math"

// Pow function retrun x^n
func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// return v // undefined: v error
	return lim
}
