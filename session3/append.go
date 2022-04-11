package main

import "fmt"

func main() {
	var s []int
	printSliceSecond(s)

	s = append(s, 0)
	printSliceSecond(s)

	s = append(s, 1)
	printSliceSecond(s)

	s = append(s, 2, 3, 4)
	printSliceSecond(s)
}

func printSliceSecond(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
