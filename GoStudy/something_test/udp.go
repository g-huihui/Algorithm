/**
 * @Author: Gong Yanhui
 * @Description: UDP测试
 * @Date: 2024-01-08 14:57
 */

package something_test

import (
	"fmt"
	"net"
)

func StartUDPListen() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("Listen failed, err: ", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, errIn := listen.ReadFromUDP(data[:]) // 接收数据
		if errIn != nil {
			fmt.Println("read udp failed, errIn: ", errIn)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, errIn = listen.WriteToUDP(data[:n], addr) // 发送数据
		if errIn != nil {
			fmt.Println("Write to udp failed, errIn: ", errIn)
			continue
		}
	}
}

func DailUDP() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接UDP服务器失败，err: ", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello Server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err: ", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败, err: ", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
