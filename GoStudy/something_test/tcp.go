/**
 * @Author: Gong Yanhui
 * @Description: TCP测试
 * @Date: 2024-01-08 15:03
 */

package something_test

import (
	"fmt"
	"net"
	"time"
)

// TestTCPNoDelay 测试TCP NoDelay字段作用 通过抓包比较开启和关闭的区别
func TestTCPNoDelay() {
	go startTCPListen()

	time.Sleep(1 * time.Second)

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("connect failed, err: ", err)
		return
	}
	defer conn.Close()
	//conn.SetNoDelay(false)
	_, err = conn.Write([]byte("hello world"))
	_, err = conn.Write([]byte("hello world"))
	_, err = conn.Write([]byte("hello world"))
	_, err = conn.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("write failed, err: ", err)
		return
	}

	time.Sleep(3 * time.Second)
}

// 开启TCP监听
func startTCPListen() {
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	for {
		conn, errIn := listen.AcceptTCP()
		if errIn != nil {
			fmt.Println("Accept failed, err: ", errIn)
			return
		}
		go handTCPConn(conn)
	}
}

func handTCPConn(conn net.Conn) {
	defer conn.Close()
	for {
		var data [1024]byte
		n, err := conn.Read(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read from connect failed, err: ", err)
			break
		}
		fmt.Println(string(data[:n]))
	}
}
