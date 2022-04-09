package main

import "fmt"

const (
	//101 << 1 → 1010(二進数)で十進数だと10です。
	//101 >> 1 → 10(二進数)で十進数だと2です。
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
