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
 *     Initial: 2018/03/26        Shi Ruitao
 */

package gostudy

import (
	"io/ioutil"
	"time"
	"strconv"
	"fmt"
	"os"
)

func getNameByTime(path string, suffix string) string {
	files, _ := ioutil.ReadDir(path)
	fmt.Println(len(files))
	for _, file:= range files {
		fmt.Println(file.Name())
	}
	timeStamp := time.Now().Unix()
	return strconv.FormatInt(timeStamp, 10) + strconv.Itoa(len(files)) + "." + suffix
}

func checkDir(path string) error {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0777)

			if err != nil {
				return err
			}
		}
	}

	return err
}

func Execute() {
	fmt.Println(getNameByTime("/users/a1/github", "a"))
	_ = checkDir("/users/a1/github/123")
}