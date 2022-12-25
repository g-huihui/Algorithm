/**
 * @Author: Gong Yanhui
 * @Description: HTTP/HTTPS 请求处理	https://laravelacademy.org/post/21040
 * @Date: 2022-12-25 17:57
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
	使用 net/http 包提供的 http.ListenAndServe() 方法，
	可以开启一个 HTTP 服务器，并且在指定的 IP 地址和端口上监听客户端请求
func ListenAndServe(addr string, handler Handler) error
*/

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		fmt.Fprintf(writer, "你好, %s", params.Get("name"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("启动 HTTP 服务器失败: %v", err)
	}
}
