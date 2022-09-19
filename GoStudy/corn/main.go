/**
 * @Author: Gong Yanhui
 * @Description: 定时任务 corn包
 * @Date: 2022-09-08 14:57
 */

package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	var count = 0
	var index = 1

	// 这种创建方式: 第三次任务还未执行完毕 第四次任务就已经开始执行了 两个任务会同时操作同一个变量 引发线程安全性问题
	//c := cron.New()

	// 触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行；
	//c := cron.New(cron.WithChain(cron.DelayIfStillRunning(cron.DefaultLogger)))

	// 触发时，如果上一次任务还未完成，则跳过此次执行。
	c := cron.New(cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))

	c.AddFunc("@every 5s", func() {
		count++
		fmt.Printf("第%v次任务", index)
		index++
		if count == 3 {
			time.Sleep(time.Second * 7)
		}
		fmt.Println(count)
	})

	c.Run()
}
