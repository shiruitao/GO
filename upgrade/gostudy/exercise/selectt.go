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
 *     Initial: 2018/04/27        Shi Ruitao
 */

package exercise

import (
	"fmt"
)

func findSmallest(arr []int) int {
	small := arr[0]
	smallIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < small {
			small = arr[i]
			smallIndex = i
		}
	}
	return smallIndex
}

func Selectt(arr []int) {
	var newArr []int
	j := len(arr)
	for i := 0; i < j; i++ {
		smallIndex := findSmallest(arr)
		newArr = append(newArr, arr[smallIndex])
		arr = append(arr[:smallIndex], arr[smallIndex+1:]...)
		//arr = arr[:smallIndex+copy(arr[smallIndex:], arr[smallIndex+1:])]
	}
	fmt.Println(newArr)
}
