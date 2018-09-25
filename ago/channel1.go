package ago

import (
	"fmt"
)

func sum(arrays []int, ch chan int) {
    //fmt.Println(arrays)
    sum := 0
    for _, array := range arrays {
        sum += array
    }
    ch <- sum
}

func main() {
    arrayChan := make(chan int, 20)
	arrayInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	length := len(arrayInt)
    for t := 0; t < 15; t++ {
        go sum(arrayInt[length - t:], arrayChan)
    }

    for i := 0; i < 15; i++ {
		fmt.Println(<-arrayChan)
    }
    
}