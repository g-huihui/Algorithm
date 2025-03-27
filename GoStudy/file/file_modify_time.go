/**
 * @Author: Gong Yanhui
 * @Description: 获取文件的修改时间
 * @Date: 2024-11-09 14:31
 */

package file

import (
	"fmt"
	"log"
	"os"
)

func GetFileModTime() {
	file, err := os.Open("/Users/gongyanhui/GolandProjects/Algorithm/GoStudy/file/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.ModTime().Unix())
	fmt.Println(fileInfo.ModTime().Format("2006-01-02 15:04:05"))
}

func GetFileModTimeSimple() {
	fileInfo, err := os.Stat("/Users/gongyanhui/GolandProjects/Algorithm/GoStudy/file/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.ModTime().Unix())
	fmt.Println(fileInfo.ModTime().Format("2006-01-02 15:04:05"))
}
