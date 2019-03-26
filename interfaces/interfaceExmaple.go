package main

import "fmt"

type Animal interface {
	Speak() string
}
type MyInt int
type MyFloat float64

type Dog struct {
}

type Cat struct {
}

type Llama struct {
}

type JavaProgrammer struct {
}

// implement methods
func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (l Llama) Speak() string {
	return "?????"
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func (mi MyInt) Speak() string {
	return "Number Number!"
}

// float MyFloat's Speak is not implemented

func main() {
	var mi MyInt = 10
	//var mf MyFloat = 10.00
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}, mi}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	var a, b Animal
	a = mi // MyInt automatically converyt to Animal type.
	//b = mf
	/*cannot use mf (type MyFloat) as type Animal in assignment:
	MyFloat does not implement Animal (missing Speak method */
	fmt.Println(b)
	//fmt.Println(b.Speak()) // invalid memory address or nil pointer dereference
	fmt.Println(a.Speak())

}
