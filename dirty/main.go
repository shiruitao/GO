package main

import (

	"bytes"
	//"strings"
	"fmt"
)

var (
	ch byte = 'A'
	name bytes.Buffer
	a = []string{"1"}
	b string
	)
func main() {
	name.WriteString(fmt.Sprintf("%d,",  1))
	println(name.String())
	b = name.String()
	println(b)
	//strings.Join(a, "1")
	//strings.Join(a, "1")
	//fmt.Println(strings.Join(a, "1"), a)
	defer println(1)
	defer println(2)
	defer println(3)

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
