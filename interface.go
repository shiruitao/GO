/*
*Shi Ruitao 2017-12-5
*/

package main
import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func main() {
	var a, b Abser

	f := MyFloat(math.Sqrt2)
	a = f
	v := Vertex{3, 4}
	b = &v

	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}