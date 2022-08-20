/**
 * @Author: Gong Yanhui
 * @Description: 1 搭个架子 实现Hello World
 * @Date: 2022-08-20 14:30
 */

package main

import (
	"fmt"
	"net"
)

/**
监听端口 1080（socks5的默认端口）
每收到一个请求，启动一个 goroutine 来处理它
*/
func main() {
	server, err := net.Listen("tcp", ":1080")
	if err != nil {
		fmt.Printf("net.Listen(\"tcp\", \":1080\") failed , err : %s\n", err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			fmt.Printf("server.Accept() failed , err : %s\n", err)
			continue
		}
		go process(client)
	}
}

func process(client net.Conn) {
	remoteAddr := client.RemoteAddr().String()

	fmt.Printf("Remote Connection from %s\n", remoteAddr)
	write, err := client.Write([]byte("Hello World!\n"))
	if err != nil {
		fmt.Printf("client.Write() failed , err : %s\n", err)
	}
	fmt.Printf("client.Write() result num is %v\n", write)

	err = client.Close()
	if err != nil {
		fmt.Printf("client.Close() failed , err : %s\n", err)
	}
}
