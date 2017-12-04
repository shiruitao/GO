/*
* Shi Ruitao  2017-12-3
*/
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}
type a struct {
	A, B, C int
}
var n = map[string]a{
	"一": a{
		1, 2, 3,
	},
	"二": a{
		12, 34, 56,
	},
}
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
	fmt.Println(n)
	m := make(map[string]int)
	
		m["Answer"] = 42
		fmt.Println("The value:", m["Answer"])
	
		m["Answer1"] = 48
		fmt.Println("The value:", m["Answer1"])
	
		delete(m, "Answer")
		fmt.Println("The value:", m["Answer"])
	
		v, ok := m["Answer1"]
		fmt.Println("The value:", v, "Present?", ok)
}