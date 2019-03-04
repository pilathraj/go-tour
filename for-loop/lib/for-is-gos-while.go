// Package lib  for is go's while
package lib

func GosWhile() int {
	sum := 1
	for sum < 100 { // condition only
		sum += sum
	}
	return sum
}
