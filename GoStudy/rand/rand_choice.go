/**
 * @Author: Gong Yanhui
 * @Description: 随机数
 * @Date: 2024-01-19 20:44
 */

package rand

import (
	"math/rand"
	"time"
)

//
// RandChoice
//  @Description: 根据随机数来做选择
//  @param count 随机次数
//  @param between 随机数的范围 根据这次选择有关系的数字来定 [0, between) 必须偶数
//
func RandChoice(count, between int) {
	var yes, no int
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	for i := 0; i < count; i++ {
		r := rand.Intn(between) % 2 // 生成随机数 0 or 1
		if r == 1 {
			yes++
		} else if r == 0 {
			no++
		} else {
			panic(r)
		}
	}
	if yes > no {
		println("yes")
	} else {
		println("no")
	}
	println(yes, no)
}
