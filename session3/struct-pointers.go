package main

import "fmt"

type VertexThird struct {
	X int
	Y int
}

func main() {
	v := VertexThird{1, 2}
	p := &v
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println((*p).X)
	p.X = 1e9
	fmt.Println(v)
}
