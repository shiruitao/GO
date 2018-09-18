/*
 * Revision History:
 *     Initial: 2018/09/03        Shi Ruitao
 */

package program

import (
	"fmt"
)

type Exercise struct {}

func Slice() {
	si1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	si2 := si1[:8]
	fmt.Println(si2)

	si2 = append(si2, 0)
	fmt.Println(si2)
	fmt.Println(si1)
}
