/*
 * Revision History:
 *     Initial: 2018/11/10        Shi Ruitao
 */

package sort

import (
	"fmt"
)

/*
 *	时间复杂度为O(n^2)
 */

var bublle = []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}

func Bubble() {
	fmt.Println("-------冒泡排序------↓")
	fmt.Println("排序前：", bublle)

	for i := 0; i < len(bublle)-1; i++ {
		for j := i + 1; j < len(bublle); j++ {
			if bublle[j] < bublle[i] {
				bublle[j], bublle[i] = bublle[i], bublle[j]
			}
		}
	}
	fmt.Println("排序后：", bublle)
}
