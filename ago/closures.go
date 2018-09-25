/*
* Shi Ruitao 2017-12-8
*/

package ago
import "fmt"

func closed() func() func() int {
	i := 0
	return func() func() int {
		i += 1
		return func() int {
			i += 1
			return i
		}
	}
}

func main() {
	fun := closed()
	// fun1 := fun()
	fmt.Println(fun()())
	fmt.Println(fun()())
	fmt.Println(fun()())
	fmt.Println(fun()())
	fmt.Println(fun()())

	fun = closed()
	fun1 := fun()
	fmt.Println(fun1())
	fmt.Println(fun1())
	fmt.Println(fun1())
	fmt.Println(fun1())
	fmt.Println(fun1())
	
}