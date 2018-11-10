/*
 * Revision History:
 *     Initial: 2018/09/25        Shi Ruitao
 */

package main

import (
	"log"
	"net/http"

	"github.com/shiruitao/GO/web"
)

func main() {
	http.HandleFunc("/aa", web.HttpTest)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
