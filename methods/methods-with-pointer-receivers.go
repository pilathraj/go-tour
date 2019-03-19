package main

import (
	"fmt"
	"math"
)

type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodsInfo() {
	v := &Vertex1{3, 4}
	fmt.Printf("Before Scaling: %v, Abs: %v\n", v, v.Abs()) // Before Scaling: &{3 4}, Abs: 5
	v.Scale(5)
	fmt.Printf("After Scaling: %v, Abs: %v\n", v, v.Abs()) //After Scaling: &{15 20}, Abs: 25
}
