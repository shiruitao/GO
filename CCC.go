/*
* Shi Ruitao  17-11-28
*/

package main
import "fmt"
func getSequence() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}

func main() {
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	
	arr := [3] int {1, 2, 3}
	fmt.Println(arr)

	var arr1 [2] int
	// arr1[0] = 5
	// arr1[1] = 6
	arr1 = [2] int {1, 2}
	fmt.Println(arr1)
}