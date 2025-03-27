/**
 * @Author: Gong Yanhui
 * @Description: 文件读取 https://houbb.github.io/2023/09/25/go-in-action-02-how-log-file-read
 * @Date: 2025-03-27 11:39
 */

package file

import (
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
)

// ReadFile Go中使用 bufio 包来逐行读取大型日志文件
func ReadFile() {
	filePath := "/Users/gongyanhui/GolandProjects/Algorithm/GoStudy/file/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // 使用该库逐行读取数据
	for scanner.Scan() {              // 从第一行开始读取
		// 获取每一行数据
		line := scanner.Text()
		processLine(line)
	}

	if errIn := scanner.Err(); errIn != nil {
		panic(errIn)
	}
	// 注意: 读取大型日志文件可能会占用较多内存 建议在读取中适当释放资源 避免内存溢出
	// 如果日志文件非常大(GB级别) 可能需要考虑分块读取或使用其他更高效的方式来处理
}

// ReadBufferFile 通过缓存区逐行读取大文件
func ReadBufferFile() {
	filePath := "/Users/gongyanhui/GolandProjects/Algorithm/GoStudy/file/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file) // 使用该库逐行读取数据

	// 定义缓冲区大小 根据实际情况调整
	bufferSize := 4096 // 确保该长度能够接收整行所有数据
	buffer := make([]byte, bufferSize)
	// 设置缓冲区
	scanner.Buffer(buffer, bufferSize)

	for scanner.Scan() {
		line := scanner.Text()
		processLine(line)

		// 如果文件一直有写入并且没有结束 scanner.Scan() 会继续读取新的行 可以自定义一些退出条件
		// if shouldTerminate() { break }
	}

	if errIn := scanner.Err(); errIn != nil {
		panic(errIn)
	}
}

func processLine(line string) {
	// 处理每一行日志数据
	// 例如 输出到控制台或者进行其他操作
	fmt.Println(line)
}

// =====何时退出文件读取=====
// 1. 根据文件指针位置: 可以通过比较文件指针的位置和文件的大小来判断文件是否读取完成
func exitReadFile1() {
	filePath := "/Users/gongyanhui/GolandProjects/Algorithm/GoStudy/file/test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, 4096) // 定义缓冲区大小 根据实际情况调整

	for {
		bytesRead, errIn := file.Read(buffer)
		if errIn != nil {
			log.Fatal(errIn)
		}
		if int64(bytesRead) == 0 {
			break // 读不到数据了 文件读取完成
		}

		processLine(string(buffer[:bytesRead]))

		if int64(bytesRead) == fileSize {
			break // 读到的数据已经是最后的数据了 文件读取完成
		}
	}
}

// 2. 根据文件变化检测: 可以使用 inotify or fsnotify 等工具来监测文件的变化 如果文件被删除、重命名、修改等操作 可以认为文件读取完成
func exitReadFile2() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	filePath := "/Users/gongyanhui/GolandProjects/Algorithm/GoStudy/file/test.txt"
	err = watcher.Add(filePath)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove || event.Op&fsnotify.Rename == fsnotify.Rename {
				// 文件已被删除或重命名 文件读取完成
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("文件被写入")
				return
			}
			// 文件发生其他变化 可以继续处理
			log.Println("continue...")
		case errIn, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("监测出错:", errIn)
		}
	}
}
