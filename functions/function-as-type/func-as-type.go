package main

import "fmt"

type CalcFunc func(int, int) int

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

//Let’s create a function which takes two integers and third argument a
//function which does mathematical operation on those two numbers.
//We will use add and subtract function as third parameter to this function.
func calc(a int, b int, f func(int, int) int) int {
	r := f(a, b)
	return r
}

// used type alias
func calc2(a int, b int, f CalcFunc) int {
	r := f(a, b)
	return r
}

func main() {
	fmt.Printf("Type of function add is %T\n", add) // func(int, int) int
	fmt.Printf("Type of function sub is %T\n", sub) //  func(int, int) int

	addResult := calc(5, 3, add)
	subResult := calc(5, 3, sub)

	fmt.Printf("5 + 3 = %d\n", addResult)
	fmt.Printf("5 - 3 = %d\n", subResult)

	fmt.Println("with Type alias:")
	addResult = calc2(5, 3, add)
	subResult = calc2(5, 3, sub)

	fmt.Printf("5 + 3 = %d\n", addResult)
	fmt.Printf("5 - 3 = %d\n", subResult)

}
