/*
* Shi Ruitao 2017-12-12
*/

package ago
import "fmt"

var i int = 1
func test() bool {
	i += 1
	return i > 1
}
func a() int {
	a := 2
	if a > 1 {
		return 0
	}
	return a
}
func main() {
	if test() {
		fmt.Println(i)
	}
	if test() {
		fmt.Println(i)
	}
	fmt.Println(a())
}