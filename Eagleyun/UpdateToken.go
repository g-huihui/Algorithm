/**
 * @Author: Gong Yanhui
 * @Description: 实现重建token覆盖旧的token
 * @Date: 2022-07-13 22:42
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var tokenFile = "/Users/gongyanhui/Downloads/token"
var tokenFile_tmp = "/Users/gongyanhui/Downloads/rename/token.tmp"

func main() {

	fmt.Println("刷新token开始...")
	for {
		// 获取token文件大小信息 确定token是否失效
		stat, err := os.Stat(tokenFile)
		if err != nil {
			fmt.Printf("os.Stat() failed: %s\n", err.Error())
			return
		} else if stat.Size() <= 1 {
			fmt.Println("token失效了...")
			return
		}

		// 读取旧token文件内容
		token, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			fmt.Printf("ioutil.ReadFile() failed: %s\n", err.Error())
			return
		}

		// 将token内容写入到token.tmp文件中
		err = ioutil.WriteFile(tokenFile_tmp, token, 0600)
		if err != nil {
			fmt.Printf("ioutil.WriteFile() failed: %s\n", err.Error())
			return
		}

		// 移动新token覆盖旧token
		err = os.Rename(tokenFile_tmp, tokenFile)
		if err != nil {
			fmt.Printf("os.Rename() failed: %s\n", err.Error())
			return
		}
	}
}
