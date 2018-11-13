/*
 * Revision History:
 *     Initial: 2018/11/10        Shi Ruitao
 */

package sort

import (
	"fmt"
	"time"
)

/*
 *	时间复杂度为O(n^2)
 */

var bublle = []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12}

func Bubble() {
	fmt.Println("-------冒泡排序------↓")
	fmt.Println("排序前：", bublle)

	t1 := time.Now()
	for i := 0; i < len(bublle)-1; i++ {
		end := false
		for j := 0; j < len(bublle)-1-i; j++ {
			if bublle[j] > bublle[j+1] {
				bublle[j], bublle[j+1] = bublle[j+1], bublle[j]
				end = true
			}
		}
		if !end {
			break
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("排序后：", bublle, elapsed)
}

func Bubble1() {
	fmt.Println("-------冒泡排序1------↓")
	fmt.Println("排序前：", bublle, bublle)

	t1 := time.Now()
	for i := 0; i < len(bublle)-1; i++ {
		for j := 0; j < len(bublle)-1-i; j++ {
			if bublle[j] > bublle[j+1] {
				bublle[j], bublle[j+1] = bublle[j+1], bublle[j]
			}
		}
	}
	elapsed := time.Since(t1)
	fmt.Println("排序后：", bublle, elapsed)
}
