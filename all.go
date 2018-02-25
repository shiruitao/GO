/*
 * MIT License
 *
 * Copyright (c) 2018 Shi Ruitao
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
 *     Initial: 2018/02/24        Shi Ruitao
 */

package main

import (
	"fmt"
	"path/filepath"
	"os"
)

var bt = []byte("shi")

func main() {
	//var b []byte
	//for i :=0 ; i < 257; i++ {
	//	b = append(b, byte(i))
	//}
	//var c uint32
	//v := int8(bt[1])
	//fmt.Println(v)
	//c = uint32(v)
	//fmt.Println(c)
	//hash32Len0to4(b)
	//xor()
	subDirToSkip := "segmentf"
	dir := "GO"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", dir, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file: %q\n", path)
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dir, err)
	}
}

func hash32Len0to4(s []byte) {
	slen := len(s)
	b := uint32(0)
	c := uint32(9)
	for i := 0; i < slen; i++ {
		v := int8(s[i])
		b = (b * 0xcc9e2d51) + uint32(v)
		//fmt.Println(b)
		c ^= b
		//fmt.Println(c)
	}
}

func xor() {
	const c1 uint32 = 0xcc9e2d51
	var a uint32 = 5
	var b uint32 = 6
	a = (a * c1) ^ (b * c1)
	//fmt.Println(a)
}