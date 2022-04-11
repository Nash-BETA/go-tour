package main

import (
	"fmt"
	"math"
)

type VertexSecond struct {
	X, Y float64
}

func Abs(v VertexSecond) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := VertexSecond{3, 4}
	fmt.Println(Abs(v))
}
