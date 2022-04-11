package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}
type Y struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func (y Y) M() {
	fmt.Println(y.S)
}

func main() {
	var i I = T{"hello"}
	var j I = Y{"おはよう"}
	i.M()
	j.M()
}
