/*
 * Revision History:
 *     Initial: 2018/11/09        Shi Ruitao
 */

package sort

import "fmt"

/*
*  	时间复杂度为O(n^2).
*/

var a = [5]int{1, 2, 4, 5, 6}
var insert = [5]int{10, 2, 4, 5, 3}
func Instertion() {
	fmt.Println("--------直接插入排序--------↓")

	var (
		c [len(a) + 1]int
		b = 3
	)
	copy(c[:], a[:])
	for i := len(a) - 1; i >= 0; i-- {
		if b > a[i] {
			c[i + 1] = b
			break
		} else {
			c[i + 1] = c[i]
		}
	}
	fmt.Println(c)
}

func InsertSort() {
	for i := 1; i < len(insert); i++ {
		t := insert[i]
		j := i - 1

		for j >= 0 {
			if t < insert[j] {
				insert[j + 1] = insert[j]
			} else {
				break
			}
			j--
		}

		insert[j + 1] = t
	}
	fmt.Println(insert)
}