/*
* Shi Ruitao  17-11-28
*/

package main
import "fmt"
func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 :%T\n",i)               
	case int:
		fmt.Println("x 是 int 型")                       
	case float64:
		fmt.Println("x 是 float64 型")           
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")                      
	case bool, string:
		fmt.Println("x 是 bool 或 string 型" )       
	default:
		fmt.Println("未知型")
	}
	arr := [5] int{11, 22, 33, 44, 55}
	for i, x := range arr {
		fmt.Println(i, x)
	}
}