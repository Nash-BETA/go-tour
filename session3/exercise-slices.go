package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	size := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		size[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			size[y][x] = uint8((x ^ y) * (x + y) / 2)
		}
	}
	return size
}

func main() {
	pic.Show(Pic)
}
