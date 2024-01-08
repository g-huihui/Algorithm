/**
 * @Author: Gong Yanhui
 * @Description: 国密监听 Server端
 * @Date: 2024-01-08 14:33
 */

package gm

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"log"
	"net/http"
	"time"
)

func Listen() {
	sigCert, err := gmtls.LoadX509KeyPair(
		"/Users/gongyanhui/Downloads/国密cert/sm2.1.sig.crt.pem",
		"/Users/gongyanhui/Downloads/国密cert/sm2.1.sig.key.pem")
	if err != nil {
		fmt.Printf("sigCert gmtls.LoadX509KeyPair failed, err:%s", err)
		return
	}
	encCert, err := gmtls.LoadX509KeyPair(
		"/Users/gongyanhui/Downloads/国密cert/sm2.1.enc.crt.pem",
		"/Users/gongyanhui/Downloads/国密cert/sm2.1.enc.key.pem")
	if err != nil {
		fmt.Printf("encCert gmtls.LoadX509KeyPair failed, err:%s", err)
		return
	}
	config := &gmtls.Config{
		GMSupport:    &gmtls.GMSupport{},
		Certificates: []gmtls.Certificate{sigCert, encCert},
	}

	listenAddr := fmt.Sprintf("0.0.0.0:%d", 8080)
	ln, err := gmtls.Listen("tcp", listenAddr, config)
	if err != nil {
		log.Fatalf("ReverseServer gmtls listen %s failed, err:%s", listenAddr, err)
		return
	}
	httpServer := &http.Server{
		ReadHeaderTimeout: 60 * time.Second, // 参考Nginx的设置
		Handler:           http.Handler(http.HandlerFunc(handleClientConn)),
	}
	err = httpServer.Serve(ln)
	if err != nil {
		log.Fatalf("Server HTTPS failed, err:%s", err)
		return
	}
}

func handleClientConn(w http.ResponseWriter, r *http.Request) {
	log.Println("handleClientConn...")
	// 读取数据打印
	_, _ = w.Write([]byte("hello world"))
}
