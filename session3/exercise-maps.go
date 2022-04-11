package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	exercise := make(map[string]int)
	for _, word := range strings.Fields(s) {
		exercise[word]++
	}
	return exercise
}

func main() {
	wc.Test(WordCount)
}
