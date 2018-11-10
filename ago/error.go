/*
*   Shi Ruitao   2017-12-6
 */
package ago

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number:" + fmt.Sprint(float64(e))
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := float64(1)
	for {
		y := z - (z*z-f)/(2*z)
		if math.Abs(y-z) < 1e-10 {
			return y, nil
		}
		z = y
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	//fmt.Println(Sqrt(-2))
}
