/*

*/

package main

import "fmt"

func main() {
    var a float32
    var b = 1.2
    c := "abc"
    d := true
    fmt.Println(a, b, c, d, "1")
    // var e, f, g = 1, 2, 3
    e, f, g := 1, 2, 3
    var (
        v1 = 11
        v2 = 22
    )
    fmt.Println(e, f, g, v1, v2)
}
