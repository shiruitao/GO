/*
* Shi Ruitao  17-11-28
*/
// x.(type)      range 用法

package ago
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
	for _, x := range arr {
		fmt.Println(x)
	}
	a:= 14.011
	d := "hello world"
	c := 10
	fmt.Printf("printf输出 %.2f %s \n", a, d)
	fmt.Println("println输出", a, d)
	fmt.Print("print输出 ", a, "\n")
	fmt.Printf("play %3d \n", c)
	var i, j int
	
	   for i=2; i < 100; i++ {
		  for j=2; j <= (i/j); j++ {
			 if(i%j==0) {
				break; // 如果发现因子，则不是素数
			 }
		  }
		  if(j > (i/j)) {
			 fmt.Printf("%d  是素数\n", i);
		  }
	   }
	   var c1, c2, c3 chan int
	   var i1, i2 int
	   select {
		  case i1 = <-c1:
			 fmt.Print("received ", i1, " from c1\n")
		  case c2 <- i2:
			 fmt.Print("sent ", i2, " to c2\n")
		  case i3, ok := (<-c3):  // same as: i3, ok := <-c3
			 if ok {
				fmt.Print("received ", i3, " from c3\n")
			 } else {
				fmt.Print("c3 is closed\n")
			 }
		  default:
			 fmt.Printf("no communication\n")
	   }
	//    a1, b1 := swap("shi", "rui")
	   a1 := "shi"
	   b1 := "rui"
	   swap(a1, b1)
	   fmt.Println(swap(a1, b1))
}
func swap(x, y string) (string, string) {
	var temp string
	temp = x
	x = y
	y = temp
	return x, y
 }