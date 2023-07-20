/**
 * @Author: Gong Yanhui
 * @Description: 实现TCPing 用于测试TCP连接的延迟
 * @Date: 2023-07-20 11:18
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	// 获取命令行参数
	pHost := flag.String("h", "", "目标主机(必传)")
	pPort := flag.String("p", "", "目标端口(必传)")
	// 获取时间间隔参数 -i
	var bTime int
	flag.IntVar(&bTime, "i", 1, "时间间隔")
	// 获取次数参数 -c
	var count int
	flag.IntVar(&count, "c", -1, "执行次数")
	flag.Parse()

	host := *pHost
	port := *pPort
	// 检查参数
	if host == "" || port == "" {
		fmt.Printf("invalid host:%s or port:%s", host, port)
		return
	}

	var tLog TcpLog
	tLog.DstPort = port
	tLog.DstHost = host

	for i := 0; ; i++ {
		if count != -1 && i >= count {
			break
		}
		beginTime := time.Now().UnixMilli()
		conn, err := net.DialTimeout("tcp", host+":"+port, 2*time.Second)
		tLog.BeginTime = beginTime
		if err != nil {
			tLog.Error = err.Error()
			fmt.Println(tLog.String())
			continue
		}
		tLog.SrcHost, tLog.SrcPort, err = net.SplitHostPort(conn.LocalAddr().String())
		if err != nil {
			tLog.Error = err.Error()
			fmt.Println(tLog.String())
			continue
		}
		_ = conn.Close()
		tLog.CostTime = time.Now().UnixMilli() - beginTime
		fmt.Println(tLog.String())
		time.Sleep(time.Duration(bTime) * time.Second)
	}
}

type TcpLog struct {
	DstHost      string `json:"dst_host"`
	DstPort      string `json:"dst_port"`
	BeginTimeStr string `json:"begin_time"`
	CostTime     int64  `json:"cost_time"`
	Error        string `json:"error"`
	SrcHost      string `json:"src_host"`
	SrcPort      string `json:"src_port"`

	BeginTime int64 `json:"-"`
}

func (t *TcpLog) String() string {
	t.BeginTimeStr = t.getTime()
	// 序列化json格式
	b, _ := json.Marshal(t)
	return string(b)
}

// 毫秒时间转换标准格式
func (t *TcpLog) getTime() string {
	return time.Unix(0, t.BeginTime*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
}
