/*
 * Revision History:
 *     Initial: 2018/08/15        Shi Ruitao
 */

package exercise

import (
	"fmt"
	//"context"
	//"golang.org/x/tools/cmd/guru/testdata/src/alias"
)

func Ctx() {
	//gen := func(ctx context.Context) <-chan int {
	//	dst := make(chan int)
	//	n := 1
	//	go func() {
	//		for {
	//			select {
	//			case <-ctx.Done():
	//				return // returning not to leak the goroutine
	//			case dst <- n:
	//				n++
	//			}
	//		}
	//	}()
	//	return dst
	//}
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // cancel when we are finished consuming integers
	//
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		break
	//	}
	//}

	fmt.Println(ysf(2, 2))
}

func ysf(n int, m int) int {
	var tmp int
	if n%m == 0 {
		tmp = n / m
	} else {
		tmp = n/m + 1
	}
	fmt.Println("人数", tmp)
	if tmp <= m+1 {
		return (tmp+1)*m - 1
	}
	path := ysf(tmp, m+1)
	return (path-2)*m + 1
}
