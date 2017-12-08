package main

import (
	"fmt"
)

type interfaceI interface {
	method()
	add()
}

type interfaceII interface {
	method()
	sub()
}

type conflict struct{}

func (this *conflict) method() {
	fmt.Println("method()")
}

func (this *conflict) sub() {
	fmt.Println("sub()")
}

func main()  {
	var (
		c	interface{}
		ok	bool
		i	interfaceI
		ii	interfaceII
	)

	c = &conflict{}
	// fmt.Println()

	//测试某个值是否实现了某个接口?
	if i, ok = c.(interfaceI); ok {
		i.method()
		fmt.Println(i, ok)
	}
	if ii, ok = c.(interfaceII); ok {
		ii.method()
		fmt.Println(ii, ok)
	}
}
