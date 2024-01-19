/**
 * @Author: Gong Yanhui
 * @Description: 单例模式
 * @Date: 2024-01-19 20:59
 */

package design_pattern

import "sync"

type Singleton struct {
	// 这是一个实例 确保只创建一次
}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
