/**
 * @Author: Gong Yanhui
 * @Description: 模拟发送icmp报文
 * @Date: 2025-03-26 21:22
 */

package ping

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

/**
ICMP报文结构 (不带首部)
| 8bit | 8bit |  16bit   |   16bit   |      16bit      | 32bit |
| Type | Code | Checksum | Identifier | Sequence Number |  Data |
*/

// Ping 实现ping功能 命令行: sudo go run main.go -w 1500 -l 28 -n 03 www.baidu.com
func Ping() {

	// 获取命令行参数
	var timeout, size, count = getCommandArgs()

	// 获取要发送icmp报文的目标地址
	var host = os.Args[len(os.Args)-1]
	//var host = "www.hao123.com"

	// 建立icmp连接
	conn, err := net.DialTimeout("ip:icmp", host, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var succeedCount, failedCount int

	fmt.Printf("PING %s (%s): %d data bytes\n", host, conn.RemoteAddr(), size)
	defer func() {
		fmt.Printf("--- %s ping statistics ---\n", host)
		fmt.Printf("%d packets transmitted, %d packets received, %.2f%% packet loss\n",
			succeedCount+failedCount, succeedCount, float32(failedCount)/float32(succeedCount+failedCount)*100)
	}()

	for i := 0; i < count; i++ {
		begin := time.Now()

		// 定义发送的数据
		data := make([]byte, size)

		// 定义 icmp回显请求 8 0
		var icmp = ICMP{Type: 8, Code: 0, ID: 1, SeqNum: 1}

		// 封装发送的数据
		var buffer bytes.Buffer
		_ = binary.Write(&buffer, binary.BigEndian, icmp) // 通过 binary库 将数据进行大端序写入
		buffer.Write(data)
		data = buffer.Bytes()

		// 生成checkSum 并将其放入data中一起发送
		cs := checkSum(data)
		data[2] = byte(cs >> 8)
		data[3] = byte(cs)

		// 设置超时时间
		_ = conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))

		// 发送数据
		_, err = conn.Write(data)
		if err != nil {
			failedCount++
			log.Println(err)
			continue
		}

		// 读响应数据
		var n int
		buf := make([]byte, 65535)
		n, err = conn.Read(buf)
		if err != nil {
			failedCount++
			log.Println(err)
			continue
		}

		fmt.Printf("%d bytes from %d.%d.%d.%d: icmp_seq=%d ttl=%d time=%d ms\n",
			n-28, buf[12], buf[13], buf[14], buf[15], i, buf[8], time.Since(begin).Milliseconds())
		time.Sleep(time.Second)
		succeedCount++
	}
}

// ICMP 发送的完整报文结构 顺序一定固定
type ICMP struct {
	Type     uint8
	Code     uint8
	Checksum uint16
	ID       uint16
	SeqNum   uint16
}

func getCommandArgs() (timeout int64, size, count int) {
	flag.Int64Var(&timeout, "w", 1000, "请求超时时长/ms")
	flag.IntVar(&size, "l", 32, "请求发送缓冲区大小/byte")
	flag.IntVar(&count, "n", 4, "发送请求数")
	flag.Parse()

	return
}

// ICMP 校验和算法
// 1. 相邻两个字节拼接到一起组成一个16bit数 将这些数累加求和
// 2. 若长度为奇数 则将剩余的1个字节也累加到求和
// 3. 得出总和后 将和值的高16位不断求和 直到高16位为0
// 4. 上述得到结果后 取反 即为校验和
func checkSum(data []byte) uint16 {
	var length = len(data)
	var index int
	var sum uint32
	for length > 1 {
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	if length > 0 { // 奇数的情况
		sum += uint32(data[index])
	}
	var high16 = sum >> 16
	for high16 != 0 {
		sum = high16 + uint32(uint16(sum))
		high16 = sum >> 16
	}
	return uint16(^sum)
}
