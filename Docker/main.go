/**
 * @Author: Gong Yanhui
 * @Description: 主要进程代码
 * @Date: 2025-05-11 11:37
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
)

// WhatIsMyIP 返回客户端IP
func WhatIsMyIP() {

	pListenPort := flag.String("p", "8080", "LISTEN PORT")
	flag.Parse()

	// 启动http服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("YOUR IP: " + r.RemoteAddr))
	})
	fmt.Println("http server start, listen port:", *pListenPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", *pListenPort), nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
	}
}

func main() {
	WhatIsMyIP()
}
