package main

import (
	"fmt"
	"math"
)

// type with non-struct type
type MyFloat float64

type Vertex struct {
	X, Y float64
}

// Abs via method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abs from func
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MyFloat's abs
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Scale method with pointer
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Scale function with pointer
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func main() {
	s := Vertex{3, 4}
	fmt.Println(s.Abs()) //5
	fmt.Println(Abs(s))  //5
	mf := MyFloat(-math.Pi)
	fmt.Println(mf)       // - 3.141592653589793
	fmt.Println(mf.Abs()) //  3.141592653589793
	mf = 5
	fmt.Println(mf.Abs()) // 5
	mf = -99
	fmt.Println(mf.Abs()) // 99

	// Re-calculate Vertex  after scale
	s.Scale(10)
	fmt.Println(s.Abs()) //50
	Scale(&s, 10)
	fmt.Println(Abs(s)) //500 Bz previous scale(s.Scale(10)) on line 57
	s1 := Vertex{4, 3}
	Scale(&s1, 10)
	fmt.Println(Abs(s1)) //50

	// Call Scale with &
	v := Vertex{4, 3}
	v.Scale(2)
	Scale(&v, 10)

	//Pointer to pass directly on Scale function
	p := &Vertex{4, 3}
	p.Scale(3)
	Scale(p, 8)
	fmt.Println(v, p) //{80 60} &{96 72}

	// pass value to the function
	q := &Vertex{4, 3}
	fmt.Println(q.Abs()) //5
	fmt.Println(Abs(*q)) // 5

	methodsInfo()

}
