// Package lib calculate sum of 100 number
package lib

func Add() int {
	sum := 1
	// init and post statements are optional.
	for sum < 1000 {
		sum += sum
	}
	return sum
}
