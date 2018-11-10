/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/02/28        Shi Ruitao
 */

package ago

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)
var chs = []chan int{ch1, ch2}

var num = []int{0, 1, 2}

func main() {
	go func() {
		fmt.Println(<-chs[0])
	}()
	go func() {
		fmt.Println(<-chs[1])
	}()
	select {
	case getCh(0) <- getNumber(1):
		fmt.Println("1th case is selected")
	case getCh(1) <- getNumber(2):
		fmt.Println("2th case is selected")
	default:
		fmt.Println("Default")
	}
	time.Sleep(1e9 * 2)
}

func getNumber(i int) int {
	fmt.Printf("num[%d]\n", i)
	return num[i]
}

func getCh(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}
