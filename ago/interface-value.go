/*
*Shi Ruitao 2017-12-5
 */
package ago

import (
	"fmt"
	//"math"
)

type I interface {
	M()
}
type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(-f)
}
func main() {
	var i I
	i = &T{"hello"}
	describe(i)
	i.M()
	i = F(2.1)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
