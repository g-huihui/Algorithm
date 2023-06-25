/**
 * @Author: Gong Yanhui
 * @Description:
 * @Date: 2023-06-20 16:39
 */

package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"log"
	"net"
	"net/http"
)

func main() {
	sigCert, err := gmtls.LoadX509KeyPair("/Users/gongyanhui/Downloads/sm2.1/sm2.1.sig.crt.pem", "/Users/gongyanhui/Downloads/sm2.1/sm2.1.sig.key.pem")
	if err != nil {
		fmt.Printf("sigCert gmtls.LoadX509KeyPair failed, err:%s", err)
		return
	}
	encCert, err := gmtls.LoadX509KeyPair("/Users/gongyanhui/Downloads/sm2.1/sm2.1.enc.crt.pem", "/Users/gongyanhui/Downloads/sm2.1/sm2.1.enc.key.pem")
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("ReverseServer accept failed, err:%s", err)
			continue
		}

		go hanldConn(conn)
	}
}

func hanldConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}

	fmt.Println(string(buf))
}

func ServerHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi, This is an example of http service in golang!\n")
	fmt.Println("Server...")
}
