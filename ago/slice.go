/*
*  Shi Ruitao  2017-12-8
*/

package ago
import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	fmt.Println(slice1)
	slice1 = []int {1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 4, 5)
	fmt.Println(slice1)
	slice2 := slice1[1:4]
	fmt.Println(slice2)
	slice2 = slice2[1:3]
	fmt.Println(slice2)
	slice1 = slice1 [:3]
	fmt.Println(slice1)
	slice1 = slice1 [:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))
}