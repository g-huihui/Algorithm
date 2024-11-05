/**
 * @Author: Gong Yanhui
 * @Description: 实现ping功能
 * @Date: 2024-11-05 10:57
 */

package go_ping

import (
	"fmt"
	"github.com/go-ping/ping"
)

func PingTest() {

	pinger, err := ping.NewPinger("www.baidu.com")
	if err != nil {
		panic(err)
	}

	pinger.Count = 3
	err = pinger.Run() // blocks until finished
	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Printf("%+v\n", stats)
}
