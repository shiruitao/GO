/*
* Shi Ruitao  17-11-28
*/

package main
import (
	"fmt"
	"time"
)
func getSequence() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}

func getAverage(arr []int, size int) float32 {
	var i, sum int

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	return float32(sum / size)
}

func main() {
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	arr := [3] int {1, 2, 3}
	fmt.Println(arr)

	var arr1 = [] int {1, 2, 13}
	var avg float32
	avg = getAverage(arr1, 3)
	fmt.Printf("avg = %f \n", avg)
	fmt.Printf("地址 %x \n", &avg)
	fmt.Println(time.Now())
}
