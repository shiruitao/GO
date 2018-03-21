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
 *     Initial: 2018/03/21        Shi Ruitao
 */

package gostudy

import (
	"sync"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

var wg = &sync.WaitGroup{}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
	wg.Done()
}

func (ua *UserAges) Get(name string) int {
	//ua.Lock()
	//defer ua.Unlock()

	defer wg.Done()
	for {
		if age, ok := ua.ages[name]; ok {
			fmt.Println(age)
			return age
		}
	}
	fmt.Println(-1)
	return -1
}

func Map() {
	wg.Add(2)
	ua := UserAges{
		make(map[string]int),
		sync.Mutex{},
	}
	go ua.Add("wang", 20)
	go ua.Get("wang")
	defer wg.Wait()
}