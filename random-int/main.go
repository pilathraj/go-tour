// package main  - generate 10 random number
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //  To actually get a random number,
	// we need to provide a unique seed for your program
	for i := 0; i <= 10; i = i + 1 {
		fmt.Println("My favorite number rand.Intn is : ", rand.Intn(50))
	}

}
