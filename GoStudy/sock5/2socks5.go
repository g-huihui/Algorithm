/**
 * @Author: Gong Yanhui
 * @Description: socks5 的架子
 * @Date: 2022-08-20 14:56
 */

package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

/**
socks5 : "SOCKS Protocol Version 5"
是个二进制协议，不那么直观，不过实际上非常简单，主要分成三个步骤：认证、建立连接、转发数据
*/
func processOn(client net.Conn) {
	defer func() {
		err := client.Close()
		if err != nil {
			fmt.Printf("client.Close() failed , err : %s\n", err)
		}
	}()

	if err := Socks5Auth(client); err != nil {
		fmt.Printf("Auth error : %s\n", err)
		return
	}

	target, err := Sock5Connect(client)
	if err != nil {
		fmt.Printf("Connect error : %s\n", err)
		return
	}

	Sock5Forward(client, target)
}

/**
socks5 协议规定，客户端需要先开口：
VER:		本次请求的协议版本号，取固定值 0x05（表示socks 5）
NMETHODS:	客户端支持的认证方式数量，可取值 1~255
METHODS: 	可用的认证方式列表
*/

// Socks5Auth 读取客户端的发言
func Socks5Auth(client net.Conn) error {
	buf := make([]byte, 256)

	// 读取 VER 和 NMETHODS
	n, err := io.ReadFull(client, buf[:2])
	if n != 2 {
		return errors.New("read header: " + err.Error())
	}

	// 判断是否是规定请求的协议版本号 - 对暗号
	ver, nMethods := int(buf[0]), int(buf[1])
	if ver != 5 {
		return errors.New("invalid version")
	}

	// 读取 METHODS 列表
	n, err = io.ReadFull(client, buf[:nMethods])
	if n != nMethods {
		return errors.New("reading methods: " + err.Error())
	}

	/**
	服务端得选择一种认证方式，告诉客户端：
	VER也是	0x05，对上 SOCKS 5 的暗号
	METHOD	选定的认证方式；其中 0x00 表示不需要认证，0x02 是用户名/密码认证，……
	*/

	// 无需认证
	n, err = client.Write([]byte{0x05, 0x00})
	if n != 2 || err != nil {
		return errors.New("write rsp err : " + err.Error())
	}

	return nil
}

/**
在完成认证以后，客户端需要告知服务端它的目标地址，协议具体要求为：
VER:		0x05，老暗号了
CMD:		连接方式，0x01=CONNECT, 0x02=BIND, 0x03=UDP ASSOCIATE
RSV:		保留字段，现在没卵用
ATYP:		地址类型，0x01=IPv4，0x03=域名，0x04=IPv6
DST.ADDR:	目标地址，细节后面讲
DST.PORT:	目标端口，2字节，网络字节序（network octec order）
*/

// Sock5Connect 先读取前四个字段
func Sock5Connect(client net.Conn) (net.Conn, error) {
	buf := make([]byte, 256)

	n, err := io.ReadFull(client, buf[:4])
	if n != 4 {
		return nil, errors.New("read header: " + err.Error())
	}

	ver, cmd, _, atyp := buf[0], buf[1], buf[2], buf[3]
	if ver != 5 || cmd != 1 {
		return nil, errors.New("invalid ver/cmd")
	}

	/**
	如前所述，ADDR 的格式取决于 ATYP：
	0x01：4个字节，对应 IPv4 地址
	0x02：先来一个字节 n 表示域名长度，然后跟着 n 个字节。注意这里不是 NUL 结尾的。
	0x03：16个字节，对应 IPv6 地址
	*/
	addr := ""
	switch atyp {
	case 1: //IP
		n, err := io.ReadFull(client, buf[:4])
		if n != 4 {
			return nil, errors.New("invalid IPv4: " + err.Error())
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])
	case 3: //域名
		n, err = io.ReadFull(client, buf[:1])
		if n != 1 {
			return nil, errors.New("invalid hostname: " + err.Error())
		}
		addrLen := int(buf[0])

		n, err = io.ReadFull(client, buf[:addrLen])
		if n != addrLen {
			return nil, errors.New("invalid hostname" + err.Error())
		}
		addr = string(buf[:addrLen])
	case 4: // IPv6 先不管
		return nil, errors.New("IPv6: no supported yet")
	default:
		return nil, errors.New("invalid atyp")
	}

	/*
		接着要读取的 PORT 是一个 2 字节的无符号整数
	*/
	n, err = io.ReadFull(client, buf[:2])
	if n != 2 {
		return nil, errors.New("read port: " + err.Error())
	}
	port := binary.BigEndian.Uint16(buf[:2])

	// 既然 ADDR 和 PORT 都就位了，我们马上创建一个到 dst 的连接：
	destAddrPort := fmt.Sprintf("%s:%d", addr, port)
	dest, err := net.Dial("tcp", destAddrPort)
	if err != nil {
		return nil, errors.New("dial dst: " + err.Error())
	}

	/**
	最后一步是告诉客户端，我们已经准备好了，协议要求是：
	VER:	暗号，还是暗号！
	REP:	状态码，0x00=成功，0x01=未知错误，……
	RSV:	依然是没卵用的 RESERVED
	ATYP:	地址类型
	BND.ADDR:	服务器和DST创建连接用的地址
	BND.PORT:	服务器和DST创建连接用的端口
	*/

	n, err = client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		dest.Close()
		return nil, errors.New("write rsp: " + err.Error())
	}
	// ATYP = 0x01 表示 IPv4，所以需要填充 6 个 0 —— 4 for ADDR, 2 for PORT。
	return dest, nil
}

/**
剩下就是转发 所谓"转发" 就是一头读 一头写
需要注意的是，由于 TCP 连接是双工通信，我们需要创建两个 goroutine，用于完成“双工转发”。
*/

func Sock5Forward(client, target net.Conn) {
	forward := func(src, dest net.Conn) {
		defer src.Close()
		defer dest.Close()
		io.Copy(src, dest)
	}
	go forward(client, target)
	go forward(target, client)
}
