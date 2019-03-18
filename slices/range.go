package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256}

func rangeInfo() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v) // 2**i
	}

	pow2 := make([]int, 10)
	for i := range pow2 { // only index
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 { //omit index with _
		fmt.Printf("%d\n", value)
	}
}
