package main

import "fmt"

type MapSecond struct {
	Lat, Long float64
}

var m2 = map[string]MapSecond{
	"Bell Labs": MapSecond{
		40.68433, -74.39967,
	},
	"Google": MapSecond{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m2)
}
