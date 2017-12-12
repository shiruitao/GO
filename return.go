/*
* Shi Ruitao 2017-12-12
*/

package main
import "fmt"

var i int = 1
func test() bool {
	i += 1
	return i > 1
}

func main() {
	fmt.Println(test())
	fmt.Println(i)
}