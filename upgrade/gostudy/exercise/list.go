package exercise

import (
	"container/list"
	"fmt"
)

var (
	i = 1
	l = list.New()
)
func List() {
	for ;i<10;i++ {
		l.PushBack(i)
	}
	fmt.Println(l.Front(), l.Front().Next())
}
