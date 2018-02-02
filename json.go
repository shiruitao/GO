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
 *     Initial: 2018/01/31        Shi Ruitao
 */

package main

import (
	"fmt"
	"encoding/json"
)

// json.Marshal() 用于将节结构转化为json
// json.Unmarshal() 将json转化为对象

type User struct {
	//结构字段必须是首字母大写
	Name  string   `json:"name"` //``中的字段用于指定json键名，若不指定，则使用结构名
	Age   int      `json:"age"`
	Sex   string   `json:"sex"`
}

func main() {
	user := &User{"tang", 20, "male"}
	fmt.Println(user)

	//将结构转化为json
	user_json, err := json.Marshal(user)
	if err != nil {
		return
	}

	//需强制转化为string类型，若不强转，则输出byte数组
	fmt.Println(string(user_json))

	//将json转回结构
	user_ := &User{}
	//第一个参数为要处理的json，第二个参数用于接收处理的结果，需指定类型
	err_ := json.Unmarshal(user_json, user_)

	if err_ != nil {
		return
	}

	fmt.Println(user_)
}
