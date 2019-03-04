// package main - print current time
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //  To actually get a random number,
	// we need to provide a unique seed for your program
	fmt.Println("My favorite number rand.Intn is : ", rand.Intn(100))

}
