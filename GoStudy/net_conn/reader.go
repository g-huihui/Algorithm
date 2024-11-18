/**
 * @Author: Gong Yanhui
 * @Description:
 * @Date: 2024-11-18 15:34
 */

package net_conn

import (
	"bufio"
	"net"
	"time"
)

var (
	Data = "hello"
	Addr = "127.0.0.1:18080"
)

func PeekData() {

	// 服务端
	go func() {
		listen, err := net.Listen("tcp", Addr)
		if err != nil {
			panic(err)
		}

		conn, errIn := listen.Accept()
		if errIn != nil {
			panic(errIn)
		}

		go handleConn(conn)
	}()

	time.Sleep(1 * time.Second)

	// 客户端
	conn, err := net.DialTimeout("tcp", Addr, 5*time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(Data))
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 查看数据 不读取
	reader := bufio.NewReader(conn)
	peek, err := reader.Peek(len(Data))
	if err != nil {
		panic(err)
	}

	println("[Server][PEEK] DATA:", string(peek))

	// 读取数据
	var buf = make([]byte, len(Data))
	_, err = reader.Read(buf)
	if err != nil {
		panic(err)
	}

	println("[Server][READ] DATA:", string(buf))
}
