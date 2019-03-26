package main

import "fmt"

func DoSomething(v interface{}) {
	fmt.Printf("%T hold %v \n", v, v)
}

type Dog struct {
}

type Strudent struct {
	id   int
	name string
}

func main() {
	DoSomething(1)           // int hold 1
	DoSomething(22.22)       // float64 hold 22.22
	DoSomething("Pilathraj") // string hold Pilathraj
	DoSomething(Dog{})       // main.Dog hold {}
	DoSomething(Strudent{})  // main.Strudent hold {0 }
}
