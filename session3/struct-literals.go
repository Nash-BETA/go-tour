package main

import "fmt"

type VertexFourth struct {
	X, Y int
}

var (
	v1 = VertexFourth{1, 2}
	v2 = VertexFourth{X: 1}
	v3 = VertexFourth{}
	p  = &VertexFourth{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
