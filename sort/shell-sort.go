/*
 * Revision History:
 *     Initial: 2018/11/10        Shi Ruitao
 */

package sort

import (
	"fmt"
)

/*
*  	时间复杂度为：
*	步长序列		最坏情况复杂度
*	 n/2^i		   O(n^2)
*	2^k - 1		   O(n^(3/2))
*	2^i * 3^j	   O(nlog^2 n)
*/

var array = []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}

func ShellSort() {
	fmt.Println("-------希尔排序-------↓")
	fmt.Println("排序前：", array)

	n := len(array)
	if n < 2 {
		return
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
			}
		}
		key = key / 2
	}

	fmt.Println("排序后：", array)
}
