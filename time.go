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

package main

import (
	"time"
	"fmt"
)

func main() {
	//t := time.NewTimer(time.Second * 3)
	////t.Reset(time.Second * 3)
	//
	//now := time.Now()
	//fmt.Printf("Now time : %v.\n", now)
	//
	//expire := <- t.C
	//fmt.Printf("Expiration time: %v.\n", expire)

	ch11 := make(chan int, 1000)
	sign := make(chan int)

	for i := 0; i < 1000; i++ {
		ch11 <- i
	}

	go func(){
		var e int
		ok := false
		var timer *time.Timer
		var t time.Time
		for{
			select {
			case e = <- ch11:
				fmt.Printf("ch11 -> %d\n",e)
			case t = <- func() <-chan time.Time {
				if timer == nil{
					//初始化到期时间据此间隔1ms的定时器
					timer = time.NewTimer(time.Millisecond)
					fmt.Println("Newtimer")
				}else {
					//复用，通过Reset方法重置定时器
					fmt.Println(time.Now())
					timer.Reset(time.Second * 2)
				}
				//得知定时器到期事件来临时，返回结果
				fmt.Println(time.Now())
				return timer.C
			}():
				fmt.Println("Timeout.", t)
				ok = true
				break
			}
			//终止for循环
			if ok {
				sign <- 0
				break
			}
		}

	}()
	<- sign
}