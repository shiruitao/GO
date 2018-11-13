/*
 * Revision History:
 *     Initial: 2018/11/09        Shi Ruitao
 */

package main

import (
	"github.com/shiruitao/GO/sort"
	"time"
)

func main() {
	//sort.Instertion()
	//sort.InsertSort()
	//
	//sort.ShellSort()

	go sort.Bubble()
	time.Sleep(time.Second)
	go sort.Bubble1()
	time.Sleep(time.Second)
}
