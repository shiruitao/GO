/*

*/

package main

import (
    "fmt"
)

func main() {
    var a float32 = 1.1
    var b = 1.2
    c := "abc"
    d := true
    fmt.Println(a, b, c, d, "1")
    // var e, f, g = 1, 2, 3
    e, f, g := 1, 2, 3
    fmt.Println(e, f, g)
    var (
        v1 = 11
        v2 = v1
    )
    const (
        v3 = iota
        v4
        v5
    )
    v2, v1 = v1, v2
    v1 --
    fmt.Println(v1, v2, v3, v4, v5)
    var ptr *int
    ptr = &v1
    fmt.Println(*ptr)

    var grade string = "B"
    var marks int = 90
 
    switch {
       case marks == 90: grade = "A"
       case marks == 80: grade = "B"
       case marks == 50, marks == 60, marks == 70 : grade = "C"
       default: grade = "D"  
    }
 
    // switch {
    //    case grade == "A" :
    //       fmt.Printf("优秀!\n" )     
    //    case grade == "B", grade == "C" :
    //       fmt.Printf("良好\n" )      
    //    case grade == "D" :
    //       fmt.Printf("及格\n" )      
    //    case grade == "F":
    //       fmt.Printf("不及格\n" )
    //    default:
    //       fmt.Printf("差\n" );
    // }
    switch grade{
    case "A" :
       fmt.Printf("优秀!\n" )     
    case "B", "C" :
       fmt.Printf("良好\n" )      
    case "D" :
       fmt.Printf("及格\n" )      
    case "F":
       fmt.Printf("不及格\n" )
    default:
       fmt.Printf("差\n" );
 }
    fmt.Printf("你的等级是 %s\n", grade );      
}
