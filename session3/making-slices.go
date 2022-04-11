package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSliceThird("a", a)

	b := make([]int, 0, 5)
	printSliceThird("b", b)

	c := b[:2]
	printSliceThird("c", c)

	d := c[2:5]
	printSliceThird("d", d)
}

func printSliceThird(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
