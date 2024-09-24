/**
 * @Author: Gong Yanhui
 * @Description: http服务 返回客户端的IP地址
 * @Date: 2024-06-17 14:24
 */

package something

import (
	"flag"
	"fmt"
	"net/http"
)

func WhatIsMyIP() {

	pListenPort := flag.String("p", "8080", "LISTEN PORT")
	flag.Parse()

	// 启动http服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("YOUR IP: " + r.RemoteAddr))
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", *pListenPort), nil)
	if err != nil {
		fmt.Println("http server start failed, err:", err)
	}
}
