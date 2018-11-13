/*
 * Revision History:
 *     Initial: 2018/11/11        Shi Ruitao
 */

package high

import (
	"log"
	"sync"

	utilruntime "github.com/fengyfei/gu/libs/util/runtime"
)

func P(works, num int, f func(interface{})) error {
	if works <= 0 || num <= 0 || f == nil {
		log.Println("err")
		return nil
	}
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)

	if num < works {
		works = num
	}
	wg := sync.WaitGroup{}
	wg.Add(works)
	for i := 0; i < works; i++ {
		go func() {
			defer utilruntime.HandleCrash()
			defer wg.Done()
			for i := range ch {
				f(i)
			}
		}()
	}
	wg.Wait()
	return nil
}