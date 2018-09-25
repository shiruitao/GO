/*
 * Revision History:
 *     Initial: 2018/09/25        Shi Ruitao
 */

package program

import (
	"strconv"
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func Str() {
	Bob := Human{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}