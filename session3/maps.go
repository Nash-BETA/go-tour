package main

import "fmt"

type MapTest struct {
	Lat, Long float64
}

var m map[string]MapTest

func main() {
	m = make(map[string]MapTest)
	m["Bell Labs"] = MapTest{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
