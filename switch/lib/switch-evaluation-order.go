// Package lib Saturday find.
package lib

import (
	"fmt"
	"time"
)

// Find when is saturday
func FindSat() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("%T as %v", today, today)
	fmt.Println(time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In 3 days")
	case today + 4:
		fmt.Println("In 4 days")
	case today + 5:
		fmt.Println("In 5 days")
	default:
		fmt.Println("Too far away.")
	}
}
