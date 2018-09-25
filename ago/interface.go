/*
*Shi Ruitao 2017-12-5
*/

package ago
import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Abser1 interface {
	Abs1() float64
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

func (v Vertex) Abs1() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
func main() {
	var a Abser
	var b Abser1
	f := MyFloat(math.Sqrt2)
	a = f
	v := Vertex{3, 4}
	b = v

	fmt.Println(a.Abs())
	fmt.Println(b.Abs1())
}