/*
* Shi Ruitao 2017-12-2
*/

package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func main() {
	v := vertex{1, 2}
	p := &v
	fmt.Println(p)
	p.X = 1e1
	fmt.Println(v)
	fmt.Println(*p)
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
