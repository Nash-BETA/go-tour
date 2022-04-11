package main

import "fmt"

type VertexSecond struct {
	X int
	Y int
}

func main() {
	v := VertexSecond{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
