package main

import (
	"fmt"
	"math"
)

type VertexThird struct {
	X, Y float64
}

func (v VertexThird) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *VertexThird) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VertexThird{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
