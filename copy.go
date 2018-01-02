package main
import (
	"fmt"
)

func main() {

var a = []int{0, 1, 2}

var s = make([]int, 6)

// 源长度为8,目标为6,只会复制前6个
n1 := copy(s, a)
fmt.Println("s - ", s)
fmt.Println("n1 - ", n1)

// 源长为7,目标为6,复制索引1到6

// 将源中的索引5,6,7复制到目标中的索引3,4,5
copy(a, a[:2])
fmt.Println(cap(a))
fmt.Println("最后", a)

}